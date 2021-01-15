package middleware

import (
	"context"
	users "inventory-gateway-service/pb"
	"log"

	"google.golang.org/grpc"
)

// Auth struct
type Auth struct {
	client users.UserServiceClient
	Token  string
}

// Unary interceptor
func (u *Auth) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		/*err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}*/

		return handler(ctx, req)
	}
}
