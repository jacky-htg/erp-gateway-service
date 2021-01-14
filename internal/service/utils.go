package service

import (
	"context"
	"inventory-gateway-service/internal/pkg/app"
	users "inventory-gateway-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getMetadataToken(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token := md["token"]
	if len(token) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	ctx = context.WithValue(ctx, app.Ctx("token"), token[0])

	return ctx, nil
}

func setMetadataToken(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{
		"token": ctx.Value(app.Ctx("token")).(string),
	})

	return metadata.NewOutgoingContext(ctx, md)
}

func setMetadata(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{
		"user_id":    ctx.Value(app.Ctx("userID")).(string),
		"company_id": ctx.Value(app.Ctx("companyID")).(string),
	})

	return metadata.NewOutgoingContext(ctx, md)
}

func getUserLoginCtx(ctx context.Context, userClient users.UserServiceClient) (context.Context, error) {
	ctx, err := getMetadataToken(ctx)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, app.Ctx("token"), ctx.Value(app.Ctx("token")).(string))
	var userLogin *users.User
	userLogin, err = userClient.GetByToken(setMetadataToken(ctx), &users.Empty{})

	ctx = context.WithValue(ctx, app.Ctx("userID"), userLogin.GetId())
	ctx = context.WithValue(ctx, app.Ctx("companyID"), userLogin.GetCompanyId())

	return ctx, nil
}
