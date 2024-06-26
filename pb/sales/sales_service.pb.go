// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: sales/sales_service.proto

package sales

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

type ListSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BranchId   string      `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	CustomerId string      `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SalesmanId string      `protobuf:"bytes,4,opt,name=salesman_id,json=salesmanId,proto3" json:"salesman_id,omitempty"`
}

func (x *ListSalesRequest) Reset() {
	*x = ListSalesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_sales_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSalesRequest) ProtoMessage() {}

func (x *ListSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_sales_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSalesRequest.ProtoReflect.Descriptor instead.
func (*ListSalesRequest) Descriptor() ([]byte, []int) {
	return file_sales_sales_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListSalesRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListSalesRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *ListSalesRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *ListSalesRequest) GetSalesmanId() string {
	if x != nil {
		return x.SalesmanId
	}
	return ""
}

type SalesPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BranchId   string      `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	CustomerId string      `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SalesmanId string      `protobuf:"bytes,4,opt,name=salesman_id,json=salesmanId,proto3" json:"salesman_id,omitempty"`
	Count      uint32      `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SalesPaginationResponse) Reset() {
	*x = SalesPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_sales_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalesPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalesPaginationResponse) ProtoMessage() {}

func (x *SalesPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_sales_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalesPaginationResponse.ProtoReflect.Descriptor instead.
func (*SalesPaginationResponse) Descriptor() ([]byte, []int) {
	return file_sales_sales_service_proto_rawDescGZIP(), []int{1}
}

func (x *SalesPaginationResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *SalesPaginationResponse) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *SalesPaginationResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *SalesPaginationResponse) GetSalesmanId() string {
	if x != nil {
		return x.SalesmanId
	}
	return ""
}

func (x *SalesPaginationResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *SalesPaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Sales      *Sales                   `protobuf:"bytes,2,opt,name=sales,proto3" json:"sales,omitempty"`
}

func (x *ListSalesResponse) Reset() {
	*x = ListSalesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_sales_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSalesResponse) ProtoMessage() {}

func (x *ListSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_sales_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSalesResponse.ProtoReflect.Descriptor instead.
func (*ListSalesResponse) Descriptor() ([]byte, []int) {
	return file_sales_sales_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSalesResponse) GetPagination() *SalesPaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListSalesResponse) GetSales() *Sales {
	if x != nil {
		return x.Sales
	}
	return nil
}

var File_sales_sales_service_proto protoreflect.FileDescriptor

var file_sales_sales_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x1a, 0x19, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x49,
	0x64, 0x22, 0xc1, 0x01, 0x0a, 0x17, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x77, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x61,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x32, 0xd4,
	0x01, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x0b, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x1a, 0x0c, 0x2e, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x73, 0x61,
	0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x1a, 0x0c, 0x2e, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x56, 0x69, 0x65, 0x77, 0x12, 0x09, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x49,
	0x64, 0x1a, 0x0c, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x35, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x0e, 0x70, 0x62,
	0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x3b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sales_sales_service_proto_rawDescOnce sync.Once
	file_sales_sales_service_proto_rawDescData = file_sales_sales_service_proto_rawDesc
)

func file_sales_sales_service_proto_rawDescGZIP() []byte {
	file_sales_sales_service_proto_rawDescOnce.Do(func() {
		file_sales_sales_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sales_sales_service_proto_rawDescData)
	})
	return file_sales_sales_service_proto_rawDescData
}

var file_sales_sales_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sales_sales_service_proto_goTypes = []interface{}{
	(*ListSalesRequest)(nil),        // 0: sales.ListSalesRequest
	(*SalesPaginationResponse)(nil), // 1: sales.SalesPaginationResponse
	(*ListSalesResponse)(nil),       // 2: sales.ListSalesResponse
	(*Pagination)(nil),              // 3: sales.Pagination
	(*Sales)(nil),                   // 4: sales.Sales
	(*Id)(nil),                      // 5: sales.Id
}
var file_sales_sales_service_proto_depIdxs = []int32{
	3, // 0: sales.ListSalesRequest.pagination:type_name -> sales.Pagination
	3, // 1: sales.SalesPaginationResponse.pagination:type_name -> sales.Pagination
	1, // 2: sales.ListSalesResponse.pagination:type_name -> sales.SalesPaginationResponse
	4, // 3: sales.ListSalesResponse.sales:type_name -> sales.Sales
	4, // 4: sales.SalesService.SalesCreate:input_type -> sales.Sales
	4, // 5: sales.SalesService.SalesUpdate:input_type -> sales.Sales
	5, // 6: sales.SalesService.SalesView:input_type -> sales.Id
	0, // 7: sales.SalesService.SalesList:input_type -> sales.ListSalesRequest
	4, // 8: sales.SalesService.SalesCreate:output_type -> sales.Sales
	4, // 9: sales.SalesService.SalesUpdate:output_type -> sales.Sales
	4, // 10: sales.SalesService.SalesView:output_type -> sales.Sales
	2, // 11: sales.SalesService.SalesList:output_type -> sales.ListSalesResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sales_sales_service_proto_init() }
func file_sales_sales_service_proto_init() {
	if File_sales_sales_service_proto != nil {
		return
	}
	file_sales_sales_message_proto_init()
	file_sales_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sales_sales_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSalesRequest); i {
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
		file_sales_sales_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalesPaginationResponse); i {
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
		file_sales_sales_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSalesResponse); i {
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
			RawDescriptor: file_sales_sales_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sales_sales_service_proto_goTypes,
		DependencyIndexes: file_sales_sales_service_proto_depIdxs,
		MessageInfos:      file_sales_sales_service_proto_msgTypes,
	}.Build()
	File_sales_sales_service_proto = out.File
	file_sales_sales_service_proto_rawDesc = nil
	file_sales_sales_service_proto_goTypes = nil
	file_sales_sales_service_proto_depIdxs = nil
}
