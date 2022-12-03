package middleware

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
)

func NewLoggingInterceptor(logger *zap.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				logger.Error("res", zap.String("procedure", req.Spec().Procedure), zap.Any("req", req.Any()), zap.Error(err))
				return nil, err
			}

			logger.Info("res", zap.String("procedure", req.Spec().Procedure), zap.Any("req", req.Any()), zap.Any("res", res.Any()))

			return res, nil
		}
	}
}
