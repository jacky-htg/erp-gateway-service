// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: inventories/receive_return_service.proto

package inventories

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListReceiveReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BranchId   string      `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ReceiveId  string      `protobuf:"bytes,3,opt,name=receive_id,json=receiveId,proto3" json:"receive_id,omitempty"`
}

func (x *ListReceiveReturnRequest) Reset() {
	*x = ListReceiveReturnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_receive_return_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReceiveReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReceiveReturnRequest) ProtoMessage() {}

func (x *ListReceiveReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_receive_return_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReceiveReturnRequest.ProtoReflect.Descriptor instead.
func (*ListReceiveReturnRequest) Descriptor() ([]byte, []int) {
	return file_inventories_receive_return_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListReceiveReturnRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListReceiveReturnRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *ListReceiveReturnRequest) GetReceiveId() string {
	if x != nil {
		return x.ReceiveId
	}
	return ""
}

type ReceiveReturnPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BranchId   string      `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ReceiveId  string      `protobuf:"bytes,3,opt,name=receive_id,json=receiveId,proto3" json:"receive_id,omitempty"`
	Count      uint32      `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ReceiveReturnPaginationResponse) Reset() {
	*x = ReceiveReturnPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_receive_return_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveReturnPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveReturnPaginationResponse) ProtoMessage() {}

func (x *ReceiveReturnPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_receive_return_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveReturnPaginationResponse.ProtoReflect.Descriptor instead.
func (*ReceiveReturnPaginationResponse) Descriptor() ([]byte, []int) {
	return file_inventories_receive_return_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveReturnPaginationResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ReceiveReturnPaginationResponse) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *ReceiveReturnPaginationResponse) GetReceiveId() string {
	if x != nil {
		return x.ReceiveId
	}
	return ""
}

func (x *ReceiveReturnPaginationResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListReceiveReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination    *ReceiveReturnPaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	ReceiveReturn *ReceiveReturn                   `protobuf:"bytes,2,opt,name=ReceiveReturn,proto3" json:"ReceiveReturn,omitempty"`
}

func (x *ListReceiveReturnResponse) Reset() {
	*x = ListReceiveReturnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_receive_return_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReceiveReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReceiveReturnResponse) ProtoMessage() {}

func (x *ListReceiveReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_receive_return_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReceiveReturnResponse.ProtoReflect.Descriptor instead.
func (*ListReceiveReturnResponse) Descriptor() ([]byte, []int) {
	return file_inventories_receive_return_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListReceiveReturnResponse) GetPagination() *ReceiveReturnPaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListReceiveReturnResponse) GetReceiveReturn() *ReceiveReturn {
	if x != nil {
		return x.ReceiveReturn
	}
	return nil
}

var File_inventories_receive_return_service_proto protoreflect.FileDescriptor

var file_inventories_receive_return_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x1a, 0x28, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x1f, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xbd, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x52, 0x0d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x32, 0xf8, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x1a, 0x23, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x1a, 0x23,
	0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x18, 0x2e,
	0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x23, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x12, 0x6b,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x47, 0x0a, 0x27, 0x63,
	0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x1a, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventories_receive_return_service_proto_rawDescOnce sync.Once
	file_inventories_receive_return_service_proto_rawDescData = file_inventories_receive_return_service_proto_rawDesc
)

func file_inventories_receive_return_service_proto_rawDescGZIP() []byte {
	file_inventories_receive_return_service_proto_rawDescOnce.Do(func() {
		file_inventories_receive_return_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventories_receive_return_service_proto_rawDescData)
	})
	return file_inventories_receive_return_service_proto_rawDescData
}

var file_inventories_receive_return_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_inventories_receive_return_service_proto_goTypes = []interface{}{
	(*ListReceiveReturnRequest)(nil),        // 0: wiradata.inventories.ListReceiveReturnRequest
	(*ReceiveReturnPaginationResponse)(nil), // 1: wiradata.inventories.ReceiveReturnPaginationResponse
	(*ListReceiveReturnResponse)(nil),       // 2: wiradata.inventories.ListReceiveReturnResponse
	(*Pagination)(nil),                      // 3: wiradata.inventories.Pagination
	(*ReceiveReturn)(nil),                   // 4: wiradata.inventories.ReceiveReturn
	(*Id)(nil),                              // 5: wiradata.inventories.Id
}
var file_inventories_receive_return_service_proto_depIdxs = []int32{
	3, // 0: wiradata.inventories.ListReceiveReturnRequest.pagination:type_name -> wiradata.inventories.Pagination
	3, // 1: wiradata.inventories.ReceiveReturnPaginationResponse.pagination:type_name -> wiradata.inventories.Pagination
	1, // 2: wiradata.inventories.ListReceiveReturnResponse.pagination:type_name -> wiradata.inventories.ReceiveReturnPaginationResponse
	4, // 3: wiradata.inventories.ListReceiveReturnResponse.ReceiveReturn:type_name -> wiradata.inventories.ReceiveReturn
	4, // 4: wiradata.inventories.ReceiveReturnService.Create:input_type -> wiradata.inventories.ReceiveReturn
	4, // 5: wiradata.inventories.ReceiveReturnService.Update:input_type -> wiradata.inventories.ReceiveReturn
	5, // 6: wiradata.inventories.ReceiveReturnService.View:input_type -> wiradata.inventories.Id
	0, // 7: wiradata.inventories.ReceiveReturnService.List:input_type -> wiradata.inventories.ListReceiveReturnRequest
	4, // 8: wiradata.inventories.ReceiveReturnService.Create:output_type -> wiradata.inventories.ReceiveReturn
	4, // 9: wiradata.inventories.ReceiveReturnService.Update:output_type -> wiradata.inventories.ReceiveReturn
	4, // 10: wiradata.inventories.ReceiveReturnService.View:output_type -> wiradata.inventories.ReceiveReturn
	2, // 11: wiradata.inventories.ReceiveReturnService.List:output_type -> wiradata.inventories.ListReceiveReturnResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_inventories_receive_return_service_proto_init() }
func file_inventories_receive_return_service_proto_init() {
	if File_inventories_receive_return_service_proto != nil {
		return
	}
	file_inventories_receive_return_message_proto_init()
	file_inventories_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inventories_receive_return_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReceiveReturnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inventories_receive_return_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveReturnPaginationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inventories_receive_return_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReceiveReturnResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventories_receive_return_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventories_receive_return_service_proto_goTypes,
		DependencyIndexes: file_inventories_receive_return_service_proto_depIdxs,
		MessageInfos:      file_inventories_receive_return_service_proto_msgTypes,
	}.Build()
	File_inventories_receive_return_service_proto = out.File
	file_inventories_receive_return_service_proto_rawDesc = nil
	file_inventories_receive_return_service_proto_goTypes = nil
	file_inventories_receive_return_service_proto_depIdxs = nil
}
