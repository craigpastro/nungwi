package middleware

import (
	"context"

	"connectrpc.com/connect"
)

type validator interface {
	Validate() error
}

func NewValidatorInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if r, ok := req.Any().(validator); ok {
				if err := r.Validate(); err != nil {
					return nil, err
				}
			}

			return next(ctx, req)
		}
	}
}
