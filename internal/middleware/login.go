package middleware

import (
	"context"
	"inventory-gateway-service/internal/pkg/app"
	users "inventory-gateway-service/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Login struct
type Login struct {
	Client users.UserServiceClient
	Token  string
}

// Unary interceptor
func (u *Login) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		if info.FullMethod != "/wiradata.users.AuthService/Login" &&
			info.FullMethod != "/wiradata.users.AuthService/ForgotPassword" &&
			info.FullMethod != "/wiradata.users.AuthService/ResetPassword" {

			// Get token from incoming metadata
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
			}

			token := md["token"]
			if len(token) == 0 {
				return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
			}

			ctx = context.WithValue(ctx, app.Ctx("token"), token[0])

			// Set token to outgoing metadata
			mdOutgoing := metadata.New(map[string]string{
				"token": ctx.Value(app.Ctx("token")).(string),
			})

			ctx = metadata.NewOutgoingContext(ctx, mdOutgoing)

			// Get User login by token
			var userLogin *users.User
			userLogin, err := u.Client.GetByToken(ctx, &users.Empty{})
			if err != nil {
				return nil, err
			}

			ctx = context.WithValue(ctx, app.Ctx("userID"), userLogin.GetId())
			ctx = context.WithValue(ctx, app.Ctx("companyID"), userLogin.GetCompanyId())

			// Set token to outgoing metadata
			mdOutgoing = metadata.New(map[string]string{
				"user_id":    ctx.Value(app.Ctx("userID")).(string),
				"company_id": ctx.Value(app.Ctx("companyID")).(string),
			})

			ctx = metadata.NewOutgoingContext(ctx, mdOutgoing)
		}
		return handler(ctx, req)
	}
}
