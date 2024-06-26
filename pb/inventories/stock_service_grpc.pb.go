// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: inventories/stock_service.proto

package inventories

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StockService_Closing_FullMethodName = "/wiradata.inventories.StockService/Closing"
	StockService_List_FullMethodName    = "/wiradata.inventories.StockService/List"
	StockService_Info_FullMethodName    = "/wiradata.inventories.StockService/Info"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	Closing(ctx context.Context, in *ClosingStockRequest, opts ...grpc.CallOption) (*MyBoolean, error)
	List(ctx context.Context, in *StockListInput, opts ...grpc.CallOption) (*StockList, error)
	Info(ctx context.Context, in *StockInfoInput, opts ...grpc.CallOption) (*StockInfo, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) Closing(ctx context.Context, in *ClosingStockRequest, opts ...grpc.CallOption) (*MyBoolean, error) {
	out := new(MyBoolean)
	err := c.cc.Invoke(ctx, StockService_Closing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) List(ctx context.Context, in *StockListInput, opts ...grpc.CallOption) (*StockList, error) {
	out := new(StockList)
	err := c.cc.Invoke(ctx, StockService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) Info(ctx context.Context, in *StockInfoInput, opts ...grpc.CallOption) (*StockInfo, error) {
	out := new(StockInfo)
	err := c.cc.Invoke(ctx, StockService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	Closing(context.Context, *ClosingStockRequest) (*MyBoolean, error)
	List(context.Context, *StockListInput) (*StockList, error)
	Info(context.Context, *StockInfoInput) (*StockInfo, error)
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) Closing(context.Context, *ClosingStockRequest) (*MyBoolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Closing not implemented")
}
func (UnimplementedStockServiceServer) List(context.Context, *StockListInput) (*StockList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStockServiceServer) Info(context.Context, *StockInfoInput) (*StockInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_Closing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosingStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Closing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_Closing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Closing(ctx, req.(*ClosingStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockListInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).List(ctx, req.(*StockListInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockInfoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Info(ctx, req.(*StockInfoInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.inventories.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Closing",
			Handler:    _StockService_Closing_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StockService_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _StockService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventories/stock_service.proto",
}
