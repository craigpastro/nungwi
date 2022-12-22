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
	"go.uber.org/zap"
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

	ctx := context.Background()
	logger := zap.Must(zap.NewDevelopment(zap.IncreaseLevel(zap.InfoLevel)))

	run(ctx, &cfg, logger)
}

func run(ctx context.Context, config *config, logger *zap.Logger) {
	if config.EnableProfiler {
		http.Handle("/debug/fgtrace", fgtrace.Config{})
		srv := &http.Server{
			Addr:              ":6060",
			ReadHeaderTimeout: 3 * time.Second,
		}

		go func() {
			logger.Info("🔬 starting fgtrace on 'localhost:6060/debug/fgtrace'")

			if err := srv.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					logger.Fatal("failed to start the profiler", zap.Error(err))
				}
			}
		}()
	}

	p := prolog.MustNew(logger)

	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(nungwiv1alphaconnect.NungwiServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle(nungwiv1alphaconnect.NewNungwiServiceHandler(
		server.NewServer(p, logger),
		connect.WithInterceptors(
			middleware.NewValidatorInterceptor(),
			middleware.NewLoggingInterceptor(logger),
		),
	))

	srv := &http.Server{
		Addr:              config.Addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		logger.Info("🙌 starting nungwi", zap.String("addr", srv.Addr))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("failed to start nungwi", zap.Error(err))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
	case <-ctx.Done():
	}
	logger.Info("nungwi attempting to shutdown gracefully")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("nungwi shutdown failed", zap.Error(err))
	}

	logger.Info("nungwi shutdown gracefully. bye 👋")
}
