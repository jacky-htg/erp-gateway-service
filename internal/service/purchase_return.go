package service

import (
	"context"
	"io"
	"log"

	"github.com/jacky-htg/erp-proto/go/pb/purchases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PurchaseReturn struct
type PurchaseReturn struct {
	Client purchases.PurchaseReturnServiceClient
	purchases.UnimplementedPurchaseReturnServiceServer
}

// Create new PurchaseReturn
func (u *PurchaseReturn) PurchaseReturnCreate(ctx context.Context, in *purchases.PurchaseReturn) (*purchases.PurchaseReturn, error) {
	return u.Client.PurchaseReturnCreate(ctx, in)
}

// Update PurchaseReturn
func (u *PurchaseReturn) PurchaseReturnUpdate(ctx context.Context, in *purchases.PurchaseReturn) (*purchases.PurchaseReturn, error) {
	return u.Client.PurchaseReturnUpdate(ctx, in)
}

// View PurchaseReturn
func (u *PurchaseReturn) PurchaseReturnView(ctx context.Context, in *purchases.Id) (*purchases.PurchaseReturn, error) {
	return u.Client.PurchaseReturnView(ctx, in)
}

// List PurchaseReturns
func (u *PurchaseReturn) PurchaseReturnList(in *purchases.ListPurchaseReturnRequest, streamServer purchases.PurchaseReturnService_PurchaseReturnListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.PurchaseReturnList(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling PurchaseReturn.List service: %s", err)
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
