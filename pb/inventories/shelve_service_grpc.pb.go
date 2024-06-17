// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: inventories/shelve_service.proto

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
	ShelveService_Create_FullMethodName = "/wiradata.inventories.ShelveService/Create"
	ShelveService_Update_FullMethodName = "/wiradata.inventories.ShelveService/Update"
	ShelveService_View_FullMethodName   = "/wiradata.inventories.ShelveService/View"
	ShelveService_Delete_FullMethodName = "/wiradata.inventories.ShelveService/Delete"
	ShelveService_List_FullMethodName   = "/wiradata.inventories.ShelveService/List"
)

// ShelveServiceClient is the client API for ShelveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShelveServiceClient interface {
	Create(ctx context.Context, in *Shelve, opts ...grpc.CallOption) (*Shelve, error)
	Update(ctx context.Context, in *Shelve, opts ...grpc.CallOption) (*Shelve, error)
	View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Shelve, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MyBoolean, error)
	List(ctx context.Context, in *ListShelveRequest, opts ...grpc.CallOption) (ShelveService_ListClient, error)
}

type shelveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShelveServiceClient(cc grpc.ClientConnInterface) ShelveServiceClient {
	return &shelveServiceClient{cc}
}

func (c *shelveServiceClient) Create(ctx context.Context, in *Shelve, opts ...grpc.CallOption) (*Shelve, error) {
	out := new(Shelve)
	err := c.cc.Invoke(ctx, ShelveService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) Update(ctx context.Context, in *Shelve, opts ...grpc.CallOption) (*Shelve, error) {
	out := new(Shelve)
	err := c.cc.Invoke(ctx, ShelveService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Shelve, error) {
	out := new(Shelve)
	err := c.cc.Invoke(ctx, ShelveService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MyBoolean, error) {
	out := new(MyBoolean)
	err := c.cc.Invoke(ctx, ShelveService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) List(ctx context.Context, in *ListShelveRequest, opts ...grpc.CallOption) (ShelveService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShelveService_ServiceDesc.Streams[0], ShelveService_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &shelveServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShelveService_ListClient interface {
	Recv() (*ListShelveResponse, error)
	grpc.ClientStream
}

type shelveServiceListClient struct {
	grpc.ClientStream
}

func (x *shelveServiceListClient) Recv() (*ListShelveResponse, error) {
	m := new(ListShelveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShelveServiceServer is the server API for ShelveService service.
// All implementations must embed UnimplementedShelveServiceServer
// for forward compatibility
type ShelveServiceServer interface {
	Create(context.Context, *Shelve) (*Shelve, error)
	Update(context.Context, *Shelve) (*Shelve, error)
	View(context.Context, *Id) (*Shelve, error)
	Delete(context.Context, *Id) (*MyBoolean, error)
	List(*ListShelveRequest, ShelveService_ListServer) error
	mustEmbedUnimplementedShelveServiceServer()
}

// UnimplementedShelveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShelveServiceServer struct {
}

func (UnimplementedShelveServiceServer) Create(context.Context, *Shelve) (*Shelve, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShelveServiceServer) Update(context.Context, *Shelve) (*Shelve, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedShelveServiceServer) View(context.Context, *Id) (*Shelve, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedShelveServiceServer) Delete(context.Context, *Id) (*MyBoolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedShelveServiceServer) List(*ListShelveRequest, ShelveService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedShelveServiceServer) mustEmbedUnimplementedShelveServiceServer() {}

// UnsafeShelveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShelveServiceServer will
// result in compilation errors.
type UnsafeShelveServiceServer interface {
	mustEmbedUnimplementedShelveServiceServer()
}

func RegisterShelveServiceServer(s grpc.ServiceRegistrar, srv ShelveServiceServer) {
	s.RegisterService(&ShelveService_ServiceDesc, srv)
}

func _ShelveService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shelve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).Create(ctx, req.(*Shelve))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shelve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).Update(ctx, req.(*Shelve))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).View(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListShelveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShelveServiceServer).List(m, &shelveServiceListServer{stream})
}

type ShelveService_ListServer interface {
	Send(*ListShelveResponse) error
	grpc.ServerStream
}

type shelveServiceListServer struct {
	grpc.ServerStream
}

func (x *shelveServiceListServer) Send(m *ListShelveResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ShelveService_ServiceDesc is the grpc.ServiceDesc for ShelveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShelveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.inventories.ShelveService",
	HandlerType: (*ShelveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShelveService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ShelveService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _ShelveService_View_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ShelveService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ShelveService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inventories/shelve_service.proto",
}
