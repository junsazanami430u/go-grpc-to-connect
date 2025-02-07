package interceptor

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
)

func NewValidateInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if msg, ok := req.Any().(interface{ Validate() error }); ok {
				if err := msg.Validate(); err != nil {
					return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("validation error: %w", err))
				}
			}

			// 次の処理を実行
			return next(ctx, req)

		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
