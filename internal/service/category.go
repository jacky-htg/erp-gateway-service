package service

import (
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/inventories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Category struct {
	Client inventories.CategoryServiceClient
	inventories.UnimplementedCategoryServiceServer
}

func (u *Category) List(in *inventories.MyEmpty, streamServer inventories.CategoryService_ListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.List(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Category.List service: %s", err)
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
