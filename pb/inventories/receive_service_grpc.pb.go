// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: inventories/receive_service.proto

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
	ReceiveService_Create_FullMethodName                = "/wiradata.inventories.ReceiveService/Create"
	ReceiveService_Update_FullMethodName                = "/wiradata.inventories.ReceiveService/Update"
	ReceiveService_View_FullMethodName                  = "/wiradata.inventories.ReceiveService/View"
	ReceiveService_List_FullMethodName                  = "/wiradata.inventories.ReceiveService/List"
	ReceiveService_OutstandingByPurchase_FullMethodName = "/wiradata.inventories.ReceiveService/OutstandingByPurchase"
)

// ReceiveServiceClient is the client API for ReceiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiveServiceClient interface {
	Create(ctx context.Context, in *Receive, opts ...grpc.CallOption) (*Receive, error)
	Update(ctx context.Context, in *Receive, opts ...grpc.CallOption) (*Receive, error)
	View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Receive, error)
	List(ctx context.Context, in *ListReceiveRequest, opts ...grpc.CallOption) (ReceiveService_ListClient, error)
	OutstandingByPurchase(ctx context.Context, in *Id, opts ...grpc.CallOption) (*OutstandingResponse, error)
}

type receiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiveServiceClient(cc grpc.ClientConnInterface) ReceiveServiceClient {
	return &receiveServiceClient{cc}
}

func (c *receiveServiceClient) Create(ctx context.Context, in *Receive, opts ...grpc.CallOption) (*Receive, error) {
	out := new(Receive)
	err := c.cc.Invoke(ctx, ReceiveService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiveServiceClient) Update(ctx context.Context, in *Receive, opts ...grpc.CallOption) (*Receive, error) {
	out := new(Receive)
	err := c.cc.Invoke(ctx, ReceiveService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiveServiceClient) View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Receive, error) {
	out := new(Receive)
	err := c.cc.Invoke(ctx, ReceiveService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiveServiceClient) List(ctx context.Context, in *ListReceiveRequest, opts ...grpc.CallOption) (ReceiveService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReceiveService_ServiceDesc.Streams[0], ReceiveService_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &receiveServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReceiveService_ListClient interface {
	Recv() (*ListReceiveResponse, error)
	grpc.ClientStream
}

type receiveServiceListClient struct {
	grpc.ClientStream
}

func (x *receiveServiceListClient) Recv() (*ListReceiveResponse, error) {
	m := new(ListReceiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *receiveServiceClient) OutstandingByPurchase(ctx context.Context, in *Id, opts ...grpc.CallOption) (*OutstandingResponse, error) {
	out := new(OutstandingResponse)
	err := c.cc.Invoke(ctx, ReceiveService_OutstandingByPurchase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiveServiceServer is the server API for ReceiveService service.
// All implementations must embed UnimplementedReceiveServiceServer
// for forward compatibility
type ReceiveServiceServer interface {
	Create(context.Context, *Receive) (*Receive, error)
	Update(context.Context, *Receive) (*Receive, error)
	View(context.Context, *Id) (*Receive, error)
	List(*ListReceiveRequest, ReceiveService_ListServer) error
	OutstandingByPurchase(context.Context, *Id) (*OutstandingResponse, error)
	mustEmbedUnimplementedReceiveServiceServer()
}

// UnimplementedReceiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReceiveServiceServer struct {
}

func (UnimplementedReceiveServiceServer) Create(context.Context, *Receive) (*Receive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReceiveServiceServer) Update(context.Context, *Receive) (*Receive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReceiveServiceServer) View(context.Context, *Id) (*Receive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedReceiveServiceServer) List(*ListReceiveRequest, ReceiveService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedReceiveServiceServer) OutstandingByPurchase(context.Context, *Id) (*OutstandingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutstandingByPurchase not implemented")
}
func (UnimplementedReceiveServiceServer) mustEmbedUnimplementedReceiveServiceServer() {}

// UnsafeReceiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiveServiceServer will
// result in compilation errors.
type UnsafeReceiveServiceServer interface {
	mustEmbedUnimplementedReceiveServiceServer()
}

func RegisterReceiveServiceServer(s grpc.ServiceRegistrar, srv ReceiveServiceServer) {
	s.RegisterService(&ReceiveService_ServiceDesc, srv)
}

func _ReceiveService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReceiveService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveServiceServer).Create(ctx, req.(*Receive))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReceiveService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReceiveService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveServiceServer).Update(ctx, req.(*Receive))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReceiveService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReceiveService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveServiceServer).View(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReceiveService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReceiveServiceServer).List(m, &receiveServiceListServer{stream})
}

type ReceiveService_ListServer interface {
	Send(*ListReceiveResponse) error
	grpc.ServerStream
}

type receiveServiceListServer struct {
	grpc.ServerStream
}

func (x *receiveServiceListServer) Send(m *ListReceiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReceiveService_OutstandingByPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveServiceServer).OutstandingByPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReceiveService_OutstandingByPurchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveServiceServer).OutstandingByPurchase(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ReceiveService_ServiceDesc is the grpc.ServiceDesc for ReceiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReceiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.inventories.ReceiveService",
	HandlerType: (*ReceiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReceiveService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReceiveService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _ReceiveService_View_Handler,
		},
		{
			MethodName: "OutstandingByPurchase",
			Handler:    _ReceiveService_OutstandingByPurchase_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ReceiveService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inventories/receive_service.proto",
}
