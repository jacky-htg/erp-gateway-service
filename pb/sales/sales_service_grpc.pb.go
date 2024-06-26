// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: sales/sales_service.proto

package sales

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
	SalesService_SalesCreate_FullMethodName = "/sales.SalesService/SalesCreate"
	SalesService_SalesUpdate_FullMethodName = "/sales.SalesService/SalesUpdate"
	SalesService_SalesView_FullMethodName   = "/sales.SalesService/SalesView"
	SalesService_SalesList_FullMethodName   = "/sales.SalesService/SalesList"
)

// SalesServiceClient is the client API for SalesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SalesServiceClient interface {
	SalesCreate(ctx context.Context, in *Sales, opts ...grpc.CallOption) (*Sales, error)
	SalesUpdate(ctx context.Context, in *Sales, opts ...grpc.CallOption) (*Sales, error)
	SalesView(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Sales, error)
	SalesList(ctx context.Context, in *ListSalesRequest, opts ...grpc.CallOption) (SalesService_SalesListClient, error)
}

type salesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSalesServiceClient(cc grpc.ClientConnInterface) SalesServiceClient {
	return &salesServiceClient{cc}
}

func (c *salesServiceClient) SalesCreate(ctx context.Context, in *Sales, opts ...grpc.CallOption) (*Sales, error) {
	out := new(Sales)
	err := c.cc.Invoke(ctx, SalesService_SalesCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesServiceClient) SalesUpdate(ctx context.Context, in *Sales, opts ...grpc.CallOption) (*Sales, error) {
	out := new(Sales)
	err := c.cc.Invoke(ctx, SalesService_SalesUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesServiceClient) SalesView(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Sales, error) {
	out := new(Sales)
	err := c.cc.Invoke(ctx, SalesService_SalesView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesServiceClient) SalesList(ctx context.Context, in *ListSalesRequest, opts ...grpc.CallOption) (SalesService_SalesListClient, error) {
	stream, err := c.cc.NewStream(ctx, &SalesService_ServiceDesc.Streams[0], SalesService_SalesList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &salesServiceSalesListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SalesService_SalesListClient interface {
	Recv() (*ListSalesResponse, error)
	grpc.ClientStream
}

type salesServiceSalesListClient struct {
	grpc.ClientStream
}

func (x *salesServiceSalesListClient) Recv() (*ListSalesResponse, error) {
	m := new(ListSalesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SalesServiceServer is the server API for SalesService service.
// All implementations must embed UnimplementedSalesServiceServer
// for forward compatibility
type SalesServiceServer interface {
	SalesCreate(context.Context, *Sales) (*Sales, error)
	SalesUpdate(context.Context, *Sales) (*Sales, error)
	SalesView(context.Context, *Id) (*Sales, error)
	SalesList(*ListSalesRequest, SalesService_SalesListServer) error
	mustEmbedUnimplementedSalesServiceServer()
}

// UnimplementedSalesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSalesServiceServer struct {
}

func (UnimplementedSalesServiceServer) SalesCreate(context.Context, *Sales) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SalesCreate not implemented")
}
func (UnimplementedSalesServiceServer) SalesUpdate(context.Context, *Sales) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SalesUpdate not implemented")
}
func (UnimplementedSalesServiceServer) SalesView(context.Context, *Id) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SalesView not implemented")
}
func (UnimplementedSalesServiceServer) SalesList(*ListSalesRequest, SalesService_SalesListServer) error {
	return status.Errorf(codes.Unimplemented, "method SalesList not implemented")
}
func (UnimplementedSalesServiceServer) mustEmbedUnimplementedSalesServiceServer() {}

// UnsafeSalesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SalesServiceServer will
// result in compilation errors.
type UnsafeSalesServiceServer interface {
	mustEmbedUnimplementedSalesServiceServer()
}

func RegisterSalesServiceServer(s grpc.ServiceRegistrar, srv SalesServiceServer) {
	s.RegisterService(&SalesService_ServiceDesc, srv)
}

func _SalesService_SalesCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sales)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).SalesCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SalesService_SalesCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).SalesCreate(ctx, req.(*Sales))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesService_SalesUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sales)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).SalesUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SalesService_SalesUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).SalesUpdate(ctx, req.(*Sales))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesService_SalesView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).SalesView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SalesService_SalesView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).SalesView(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesService_SalesList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSalesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SalesServiceServer).SalesList(m, &salesServiceSalesListServer{stream})
}

type SalesService_SalesListServer interface {
	Send(*ListSalesResponse) error
	grpc.ServerStream
}

type salesServiceSalesListServer struct {
	grpc.ServerStream
}

func (x *salesServiceSalesListServer) Send(m *ListSalesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SalesService_ServiceDesc is the grpc.ServiceDesc for SalesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SalesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sales.SalesService",
	HandlerType: (*SalesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SalesCreate",
			Handler:    _SalesService_SalesCreate_Handler,
		},
		{
			MethodName: "SalesUpdate",
			Handler:    _SalesService_SalesUpdate_Handler,
		},
		{
			MethodName: "SalesView",
			Handler:    _SalesService_SalesView_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SalesList",
			Handler:       _SalesService_SalesList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sales/sales_service.proto",
}
