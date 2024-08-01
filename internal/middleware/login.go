package middleware

import (
	"context"
	"log"

	"github.com/jacky-htg/erp-pkg/app"
	"github.com/jacky-htg/erp-proto/go/pb/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Login struct
type Login struct {
	Client users.UserServiceClient
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

// Context method to override the context
func (w *wrappedStream) Context() context.Context {
	return w.ctx
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
		var err error

		if info.FullMethod != "/wiradata.users.AuthService/Login" &&
			info.FullMethod != "/wiradata.users.AuthService/ForgotPassword" &&
			info.FullMethod != "/wiradata.users.AuthService/ResetPassword" &&
			info.FullMethod != "/wiradata.users.CompanyService/Registration" {
			ctx, err = u.login(ctx)
			if err != nil {
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}

// Stream interceptor
func (u *Login) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx, err := u.login(stream.Context())
		if err != nil {
			return err
		}

		wrappedStream := &wrappedStream{
			ServerStream: stream,
			ctx:          ctx,
		}

		return handler(srv, wrappedStream)
	}
}

func (u *Login) login(ctx context.Context) (context.Context, error) {
	// Get token from incoming metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token := md["token"]
	if len(token) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	ctx = context.WithValue(ctx, app.Ctx("token"), token[0])

	// Set token to outgoing metadata
	mdOutgoing := metadata.New(map[string]string{
		"token": ctx.Value(app.Ctx("token")).(string),
	})

	ctx = metadata.NewOutgoingContext(ctx, mdOutgoing)

	// Get User login by token
	var userLogin *users.User
	userLogin, err := u.Client.GetByToken(ctx, &users.MyEmpty{})
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, app.Ctx("userID"), userLogin.GetId())
	ctx = context.WithValue(ctx, app.Ctx("companyID"), userLogin.GetCompanyId())

	// Set token to outgoing metadata
	mdOutgoing = metadata.New(map[string]string{
		"user_id":    ctx.Value(app.Ctx("userID")).(string),
		"company_id": ctx.Value(app.Ctx("companyID")).(string),
	})

	ctx = metadata.NewOutgoingContext(ctx, mdOutgoing)

	return ctx, nil
}
