// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: users/generic_message.proto

package users

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

type Pagination_Sort int32

const (
	Pagination_ASC  Pagination_Sort = 0
	Pagination_DESC Pagination_Sort = 1
)

// Enum value maps for Pagination_Sort.
var (
	Pagination_Sort_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	Pagination_Sort_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x Pagination_Sort) Enum() *Pagination_Sort {
	p := new(Pagination_Sort)
	*p = x
	return p
}

func (x Pagination_Sort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pagination_Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_users_generic_message_proto_enumTypes[0].Descriptor()
}

func (Pagination_Sort) Type() protoreflect.EnumType {
	return &file_users_generic_message_proto_enumTypes[0]
}

func (x Pagination_Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pagination_Sort.Descriptor instead.
func (Pagination_Sort) EnumDescriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{5, 0}
}

type MyEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MyEmpty) Reset() {
	*x = MyEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_generic_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyEmpty) ProtoMessage() {}

func (x *MyEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_users_generic_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyEmpty.ProtoReflect.Descriptor instead.
func (*MyEmpty) Descriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{0}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_generic_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_users_generic_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{1}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MyString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *MyString) Reset() {
	*x = MyString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_generic_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyString) ProtoMessage() {}

func (x *MyString) ProtoReflect() protoreflect.Message {
	mi := &file_users_generic_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyString.ProtoReflect.Descriptor instead.
func (*MyString) Descriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{2}
}

func (x *MyString) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_generic_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_users_generic_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MyBoolean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boolean bool `protobuf:"varint,1,opt,name=boolean,proto3" json:"boolean,omitempty"`
}

func (x *MyBoolean) Reset() {
	*x = MyBoolean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_generic_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyBoolean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyBoolean) ProtoMessage() {}

func (x *MyBoolean) ProtoReflect() protoreflect.Message {
	mi := &file_users_generic_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyBoolean.ProtoReflect.Descriptor instead.
func (*MyBoolean) Descriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{4}
}

func (x *MyBoolean) GetBoolean() bool {
	if x != nil {
		return x.Boolean
	}
	return false
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    uint32          `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit   uint32          `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  uint32          `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Search  string          `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
	OrderBy string          `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Sort    Pagination_Sort `protobuf:"varint,6,opt,name=sort,proto3,enum=wiradata.users.Pagination_Sort" json:"sort,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_generic_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_users_generic_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_users_generic_message_proto_rawDescGZIP(), []int{5}
}

func (x *Pagination) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Pagination) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *Pagination) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *Pagination) GetSort() Pagination_Sort {
	if x != nil {
		return x.Sort
	}
	return Pagination_ASC
}

var File_users_generic_message_proto protoreflect.FileDescriptor

var file_users_generic_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x09, 0x0a,
	0x07, 0x4d, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22,
	0x0a, 0x08, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x22, 0xd1,
	0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x19, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43,
	0x10, 0x01, 0x42, 0x35, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_users_generic_message_proto_rawDescOnce sync.Once
	file_users_generic_message_proto_rawDescData = file_users_generic_message_proto_rawDesc
)

func file_users_generic_message_proto_rawDescGZIP() []byte {
	file_users_generic_message_proto_rawDescOnce.Do(func() {
		file_users_generic_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_generic_message_proto_rawDescData)
	})
	return file_users_generic_message_proto_rawDescData
}

var file_users_generic_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_users_generic_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_users_generic_message_proto_goTypes = []interface{}{
	(Pagination_Sort)(0), // 0: wiradata.users.Pagination.Sort
	(*MyEmpty)(nil),      // 1: wiradata.users.MyEmpty
	(*Id)(nil),           // 2: wiradata.users.Id
	(*MyString)(nil),     // 3: wiradata.users.MyString
	(*Message)(nil),      // 4: wiradata.users.Message
	(*MyBoolean)(nil),    // 5: wiradata.users.MyBoolean
	(*Pagination)(nil),   // 6: wiradata.users.Pagination
}
var file_users_generic_message_proto_depIdxs = []int32{
	0, // 0: wiradata.users.Pagination.sort:type_name -> wiradata.users.Pagination.Sort
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_users_generic_message_proto_init() }
func file_users_generic_message_proto_init() {
	if File_users_generic_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_generic_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyEmpty); i {
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
		file_users_generic_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_users_generic_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyString); i {
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
		file_users_generic_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_users_generic_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyBoolean); i {
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
		file_users_generic_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_users_generic_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_users_generic_message_proto_goTypes,
		DependencyIndexes: file_users_generic_message_proto_depIdxs,
		EnumInfos:         file_users_generic_message_proto_enumTypes,
		MessageInfos:      file_users_generic_message_proto_msgTypes,
	}.Build()
	File_users_generic_message_proto = out.File
	file_users_generic_message_proto_rawDesc = nil
	file_users_generic_message_proto_goTypes = nil
	file_users_generic_message_proto_depIdxs = nil
}
