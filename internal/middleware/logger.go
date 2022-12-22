package middleware

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
)

func NewLoggingInterceptor(logger *zap.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			res, err := next(ctx, req)

			took := time.Since(start)
			fields := []zap.Field{
				zap.Duration("took", took),
				zap.String("procedure", req.Spec().Procedure),
				zap.Any("req", req.Any()),
			}

			if err != nil {
				fields = append(fields, zap.Error(err))
				logger.Error("rpc_error", fields...)
				return nil, err
			}

			fields = append(fields, zap.Any("res", res.Any()))
			logger.Debug("rpc_complete", fields...)

			res.Header().Add("X-Response-Time", took.String())

			return res, nil
		}
	}
}
