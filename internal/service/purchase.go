package service

import (
	"context"
	"erp-gateway/pb/purchases"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Purchase struct
type Purchase struct {
	Client purchases.PurchaseServiceClient
	purchases.UnimplementedPurchaseServiceServer
}

// Create new Purchase
func (u *Purchase) PurchaseCreate(ctx context.Context, in *purchases.Purchase) (*purchases.Purchase, error) {
	return u.Client.PurchaseCreate(ctx, in)
}

// Update Purchase
func (u *Purchase) PurchaseUpdate(ctx context.Context, in *purchases.Purchase) (*purchases.Purchase, error) {
	return u.Client.PurchaseUpdate(ctx, in)
}

// View Purchase
func (u *Purchase) PurchaseView(ctx context.Context, in *purchases.Id) (*purchases.Purchase, error) {
	return u.Client.PurchaseView(ctx, in)
}

// List Purchases
func (u *Purchase) PurchaseList(in *purchases.ListPurchaseRequest, streamServer purchases.PurchaseService_PurchaseListServer) error {
	ctx := streamServer.Context()
	streamClient, err := u.Client.PurchaseList(ctx, in)
	if err != nil {
		return status.Errorf(codes.Internal, "Error when calling Purchase.List service: %s", err)
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
