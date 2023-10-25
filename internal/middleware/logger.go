package middleware

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/exp/slog"
)

func NewLoggingInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			res, err := next(ctx, req)

			took := slog.Duration("took", time.Since(start))
			procedure := slog.String("procedure", req.Spec().Procedure)
			requ := slog.Any("req", req.Any())

			if err != nil {
				slog.Error("rpc_error", err, took, procedure, requ)
				return nil, err
			}

			slog.Debug("rpc_complete", took, procedure, requ, slog.Any("res", res.Any()))

			res.Header().Add("X-Response-Time", took.String())

			return res, nil
		}
	}
}
