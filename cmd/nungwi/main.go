package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha/nungwiv1alphaconnect"
	"github.com/craigpastro/nungwi/internal/middleware"
	"github.com/craigpastro/nungwi/internal/prolog"
	"github.com/craigpastro/nungwi/server"
	"github.com/felixge/fgtrace"
	"github.com/sethvargo/go-envconfig"
	"golang.org/x/exp/slog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type config struct {
	Addr           string `env:"ADDR,default=:8080"`
	EnableProfiler bool   `env:"ENABLE_PROFILER,default=false"`
}

func main() {
	var cfg config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		log.Fatal(err)
	}

	handler := slog.NewTextHandler(os.Stdout)
	logger := slog.New(handler)
	ctx := slog.NewContext(context.Background(), logger)

	run(ctx, &cfg)
}

func run(ctx context.Context, config *config) {
	if config.EnableProfiler {
		http.Handle("/debug/fgtrace", fgtrace.Config{})
		srv := &http.Server{
			Addr:              ":6060",
			ReadHeaderTimeout: 3 * time.Second,
		}

		go func() {
			slog.Info("🔬 starting fgtrace on 'localhost:6060/debug/fgtrace'")

			if err := srv.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					slog.Error("failed to start the profiler", err)
					os.Exit(1)
				}
			}
		}()
	}

	p := prolog.MustNew()

	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(nungwiv1alphaconnect.NungwiServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle(nungwiv1alphaconnect.NewNungwiServiceHandler(
		server.NewServer(p),
		connect.WithInterceptors(
			middleware.NewValidatorInterceptor(),
			middleware.NewLoggingInterceptor(),
		),
	))

	srv := &http.Server{
		Addr:              config.Addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		slog.Info("🙌 starting nungwi", slog.String("addr", srv.Addr))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start nungwi", err)
			os.Exit(1)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
	case <-ctx.Done():
	}
	slog.Info("nungwi attempting to shutdown gracefully")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("nungwi shutdown failed", err)
		os.Exit(1)
	}

	slog.Info("nungwi shutdown gracefully. bye 👋")
}
