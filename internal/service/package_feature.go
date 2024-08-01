package service

import (
	"context"
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PackageFeature struct
type PackageFeature struct {
	Client users.PackageFeatureServiceClient
	users.UnimplementedPackageFeatureServiceServer
}

// View Package Feature
func (u *PackageFeature) View(ctx context.Context, in *users.Id) (*users.PackageOfFeature, error) {
	return u.Client.View(ctx, in)
}

// List PackageFeature
func (u *PackageFeature) List(in *users.MyEmpty, streamServer users.PackageFeatureService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling User.List service: %s", err)
	}

	for {
		resp, err := streamClient.Recv()
		if err == io.EOF {
			log.Print("end stream")
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "cannot receive %v", err)
		}

		err = streamServer.Send(resp)
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
		}
	}

	return nil
}
