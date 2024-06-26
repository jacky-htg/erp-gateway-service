// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: inventories/receive_return_detail_message.proto

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

type ReceiveReturnDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReceiveReturnId string   `protobuf:"bytes,2,opt,name=receive_return_id,json=receiveReturnId,proto3" json:"receive_return_id,omitempty"`
	Product         *Product `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	Shelve          *Shelve  `protobuf:"bytes,4,opt,name=shelve,proto3" json:"shelve,omitempty"`
	Code            string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Qty             uint32   `protobuf:"varint,6,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *ReceiveReturnDetail) Reset() {
	*x = ReceiveReturnDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_receive_return_detail_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveReturnDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveReturnDetail) ProtoMessage() {}

func (x *ReceiveReturnDetail) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_receive_return_detail_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveReturnDetail.ProtoReflect.Descriptor instead.
func (*ReceiveReturnDetail) Descriptor() ([]byte, []int) {
	return file_inventories_receive_return_detail_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveReturnDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReceiveReturnDetail) GetReceiveReturnId() string {
	if x != nil {
		return x.ReceiveReturnId
	}
	return ""
}

func (x *ReceiveReturnDetail) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ReceiveReturnDetail) GetShelve() *Shelve {
	if x != nil {
		return x.Shelve
	}
	return nil
}

func (x *ReceiveReturnDetail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ReceiveReturnDetail) GetQty() uint32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

var File_inventories_receive_return_detail_message_proto protoreflect.FileDescriptor

var file_inventories_receive_return_detail_message_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x21, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x52, 0x06, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x71, 0x74, 0x79, 0x42, 0x47, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x50, 0x01, 0x5a, 0x1a, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventories_receive_return_detail_message_proto_rawDescOnce sync.Once
	file_inventories_receive_return_detail_message_proto_rawDescData = file_inventories_receive_return_detail_message_proto_rawDesc
)

func file_inventories_receive_return_detail_message_proto_rawDescGZIP() []byte {
	file_inventories_receive_return_detail_message_proto_rawDescOnce.Do(func() {
		file_inventories_receive_return_detail_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventories_receive_return_detail_message_proto_rawDescData)
	})
	return file_inventories_receive_return_detail_message_proto_rawDescData
}

var file_inventories_receive_return_detail_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_inventories_receive_return_detail_message_proto_goTypes = []interface{}{
	(*ReceiveReturnDetail)(nil), // 0: wiradata.inventories.ReceiveReturnDetail
	(*Product)(nil),             // 1: wiradata.inventories.Product
	(*Shelve)(nil),              // 2: wiradata.inventories.Shelve
}
var file_inventories_receive_return_detail_message_proto_depIdxs = []int32{
	1, // 0: wiradata.inventories.ReceiveReturnDetail.product:type_name -> wiradata.inventories.Product
	2, // 1: wiradata.inventories.ReceiveReturnDetail.shelve:type_name -> wiradata.inventories.Shelve
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_inventories_receive_return_detail_message_proto_init() }
func file_inventories_receive_return_detail_message_proto_init() {
	if File_inventories_receive_return_detail_message_proto != nil {
		return
	}
	file_inventories_product_message_proto_init()
	file_inventories_shelve_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inventories_receive_return_detail_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveReturnDetail); i {
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
			RawDescriptor: file_inventories_receive_return_detail_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventories_receive_return_detail_message_proto_goTypes,
		DependencyIndexes: file_inventories_receive_return_detail_message_proto_depIdxs,
		MessageInfos:      file_inventories_receive_return_detail_message_proto_msgTypes,
	}.Build()
	File_inventories_receive_return_detail_message_proto = out.File
	file_inventories_receive_return_detail_message_proto_rawDesc = nil
	file_inventories_receive_return_detail_message_proto_goTypes = nil
	file_inventories_receive_return_detail_message_proto_depIdxs = nil
}
