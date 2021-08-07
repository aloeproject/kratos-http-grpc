package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

type HandlerFunc func(ctx context.Context, req, err interface{}) error

type options struct {
	handler HandlerFunc
}

func CheckHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			options := options{
				handler: func(ctx context.Context, req, err interface{}) error {
					return errors.Newf(400,"", "recovery", "haha")
				},
			}

			err = options.handler(ctx, req, nil)
			return nil, err
		}
	}
}
