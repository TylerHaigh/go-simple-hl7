// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: hl7-server.proto

package proto

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

type HL7Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HL7Message) Reset() {
	*x = HL7Message{}
	mi := &file_hl7_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HL7Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HL7Message) ProtoMessage() {}

func (x *HL7Message) ProtoReflect() protoreflect.Message {
	mi := &file_hl7_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HL7Message.ProtoReflect.Descriptor instead.
func (*HL7Message) Descriptor() ([]byte, []int) {
	return file_hl7_server_proto_rawDescGZIP(), []int{0}
}

func (x *HL7Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_hl7_server_proto protoreflect.FileDescriptor

var file_hl7_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x6c, 0x37, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x68, 0x6c, 0x37, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x26, 0x0a,
	0x0a, 0x48, 0x4c, 0x37, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x49, 0x0a, 0x0a, 0x48, 0x4c, 0x37, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x15, 0x2e, 0x68, 0x6c, 0x37, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48,
	0x4c, 0x37, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x68, 0x6c, 0x37, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x4c, 0x37, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54,
	0x79, 0x6c, 0x65, 0x72, 0x48, 0x61, 0x69, 0x67, 0x68, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2d, 0x68, 0x6c, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hl7_server_proto_rawDescOnce sync.Once
	file_hl7_server_proto_rawDescData = file_hl7_server_proto_rawDesc
)

func file_hl7_server_proto_rawDescGZIP() []byte {
	file_hl7_server_proto_rawDescOnce.Do(func() {
		file_hl7_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_hl7_server_proto_rawDescData)
	})
	return file_hl7_server_proto_rawDescData
}

var file_hl7_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hl7_server_proto_goTypes = []any{
	(*HL7Message)(nil), // 0: hl7Server.HL7Message
}
var file_hl7_server_proto_depIdxs = []int32{
	0, // 0: hl7Server.HL7Service.SendMessage:input_type -> hl7Server.HL7Message
	0, // 1: hl7Server.HL7Service.SendMessage:output_type -> hl7Server.HL7Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hl7_server_proto_init() }
func file_hl7_server_proto_init() {
	if File_hl7_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hl7_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hl7_server_proto_goTypes,
		DependencyIndexes: file_hl7_server_proto_depIdxs,
		MessageInfos:      file_hl7_server_proto_msgTypes,
	}.Build()
	File_hl7_server_proto = out.File
	file_hl7_server_proto_rawDesc = nil
	file_hl7_server_proto_goTypes = nil
	file_hl7_server_proto_depIdxs = nil
}