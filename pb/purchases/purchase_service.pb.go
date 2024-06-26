// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: purchases/purchase_service.proto

package purchases

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

type ListPurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BranchId   string      `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	SupplierId string      `protobuf:"bytes,3,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
}

func (x *ListPurchaseRequest) Reset() {
	*x = ListPurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchases_purchase_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPurchaseRequest) ProtoMessage() {}

func (x *ListPurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_purchases_purchase_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPurchaseRequest.ProtoReflect.Descriptor instead.
func (*ListPurchaseRequest) Descriptor() ([]byte, []int) {
	return file_purchases_purchase_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListPurchaseRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPurchaseRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *ListPurchaseRequest) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

type PurchasePaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BranchId   string      `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	SupplierId string      `protobuf:"bytes,3,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	Count      uint32      `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PurchasePaginationResponse) Reset() {
	*x = PurchasePaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchases_purchase_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchasePaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchasePaginationResponse) ProtoMessage() {}

func (x *PurchasePaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_purchases_purchase_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchasePaginationResponse.ProtoReflect.Descriptor instead.
func (*PurchasePaginationResponse) Descriptor() ([]byte, []int) {
	return file_purchases_purchase_service_proto_rawDescGZIP(), []int{1}
}

func (x *PurchasePaginationResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *PurchasePaginationResponse) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *PurchasePaginationResponse) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *PurchasePaginationResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListPurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *PurchasePaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Purchase   *Purchase                   `protobuf:"bytes,2,opt,name=purchase,proto3" json:"purchase,omitempty"`
}

func (x *ListPurchaseResponse) Reset() {
	*x = ListPurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchases_purchase_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPurchaseResponse) ProtoMessage() {}

func (x *ListPurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_purchases_purchase_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPurchaseResponse.ProtoReflect.Descriptor instead.
func (*ListPurchaseResponse) Descriptor() ([]byte, []int) {
	return file_purchases_purchase_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPurchaseResponse) GetPagination() *PurchasePaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPurchaseResponse) GetPurchase() *Purchase {
	if x != nil {
		return x.Purchase
	}
	return nil
}

var File_purchases_purchase_service_proto protoreflect.FileDescriptor

var file_purchases_purchase_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x20, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa7, 0x01,
	0x0a, 0x1a, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x32, 0x89, 0x03, 0x0a, 0x0f, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0e,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x1a, 0x13, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0d, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53,
	0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e,
	0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x6f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x00, 0x42, 0x41, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x50, 0x01, 0x5a,
	0x16, 0x70, 0x62, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x3b, 0x70, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchases_purchase_service_proto_rawDescOnce sync.Once
	file_purchases_purchase_service_proto_rawDescData = file_purchases_purchase_service_proto_rawDesc
)

func file_purchases_purchase_service_proto_rawDescGZIP() []byte {
	file_purchases_purchase_service_proto_rawDescOnce.Do(func() {
		file_purchases_purchase_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchases_purchase_service_proto_rawDescData)
	})
	return file_purchases_purchase_service_proto_rawDescData
}

var file_purchases_purchase_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_purchases_purchase_service_proto_goTypes = []interface{}{
	(*ListPurchaseRequest)(nil),        // 0: purchases.ListPurchaseRequest
	(*PurchasePaginationResponse)(nil), // 1: purchases.PurchasePaginationResponse
	(*ListPurchaseResponse)(nil),       // 2: purchases.ListPurchaseResponse
	(*Pagination)(nil),                 // 3: purchases.Pagination
	(*Purchase)(nil),                   // 4: purchases.Purchase
	(*Id)(nil),                         // 5: purchases.Id
	(*OutstandingPurchaseRequest)(nil), // 6: purchases.OutstandingPurchaseRequest
	(*OutstandingPurchaseDetails)(nil), // 7: purchases.OutstandingPurchaseDetails
}
var file_purchases_purchase_service_proto_depIdxs = []int32{
	3, // 0: purchases.ListPurchaseRequest.pagination:type_name -> purchases.Pagination
	3, // 1: purchases.PurchasePaginationResponse.pagination:type_name -> purchases.Pagination
	1, // 2: purchases.ListPurchaseResponse.pagination:type_name -> purchases.PurchasePaginationResponse
	4, // 3: purchases.ListPurchaseResponse.purchase:type_name -> purchases.Purchase
	4, // 4: purchases.PurchaseService.PurchaseCreate:input_type -> purchases.Purchase
	4, // 5: purchases.PurchaseService.PurchaseUpdate:input_type -> purchases.Purchase
	5, // 6: purchases.PurchaseService.PurchaseView:input_type -> purchases.Id
	0, // 7: purchases.PurchaseService.PurchaseList:input_type -> purchases.ListPurchaseRequest
	6, // 8: purchases.PurchaseService.GetOutstandingPurchaseDetails:input_type -> purchases.OutstandingPurchaseRequest
	4, // 9: purchases.PurchaseService.PurchaseCreate:output_type -> purchases.Purchase
	4, // 10: purchases.PurchaseService.PurchaseUpdate:output_type -> purchases.Purchase
	4, // 11: purchases.PurchaseService.PurchaseView:output_type -> purchases.Purchase
	2, // 12: purchases.PurchaseService.PurchaseList:output_type -> purchases.ListPurchaseResponse
	7, // 13: purchases.PurchaseService.GetOutstandingPurchaseDetails:output_type -> purchases.OutstandingPurchaseDetails
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_purchases_purchase_service_proto_init() }
func file_purchases_purchase_service_proto_init() {
	if File_purchases_purchase_service_proto != nil {
		return
	}
	file_purchases_purchase_message_proto_init()
	file_purchases_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_purchases_purchase_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPurchaseRequest); i {
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
		file_purchases_purchase_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchasePaginationResponse); i {
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
		file_purchases_purchase_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPurchaseResponse); i {
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
			RawDescriptor: file_purchases_purchase_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_purchases_purchase_service_proto_goTypes,
		DependencyIndexes: file_purchases_purchase_service_proto_depIdxs,
		MessageInfos:      file_purchases_purchase_service_proto_msgTypes,
	}.Build()
	File_purchases_purchase_service_proto = out.File
	file_purchases_purchase_service_proto_rawDesc = nil
	file_purchases_purchase_service_proto_goTypes = nil
	file_purchases_purchase_service_proto_depIdxs = nil
}
