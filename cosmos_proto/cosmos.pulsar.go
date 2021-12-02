package cosmos_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: cosmos_proto/cosmos.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_cosmos_proto_cosmos_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93001,
		Name:          "cosmos_proto.implements_interface",
		Tag:           "bytes,93001,opt,name=implements_interface",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93001,
		Name:          "cosmos_proto.accepts_interface",
		Tag:           "bytes,93001,opt,name=accepts_interface",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93002,
		Name:          "cosmos_proto.scalar",
		Tag:           "bytes,93002,opt,name=scalar",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93003,
		Name:          "cosmos_proto.go_type",
		Tag:           "bytes,93003,opt,name=go_type",
		Filename:      "cosmos_proto/cosmos.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// implements_interface is used to annotate interface implementations
	//
	// optional string implements_interface = 93001;
	E_ImplementsInterface = &file_cosmos_proto_cosmos_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// accepts_interface is used to annote fields that accept interfaces
	//
	// optional string accepts_interface = 93001;
	E_AcceptsInterface = &file_cosmos_proto_cosmos_proto_extTypes[1]
	// scalar is used to define scalar types
	//
	// optional string scalar = 93002;
	E_Scalar = &file_cosmos_proto_cosmos_proto_extTypes[2]
	// go_type defines a custom go type for the field.
	// NOTE: it is not valid for repeated and map fields.
	//
	// optional string go_type = 93003;
	E_GoType = &file_cosmos_proto_cosmos_proto_extTypes[3]
)

var File_cosmos_proto_cosmos_proto protoreflect.FileDescriptor

var file_cosmos_proto_cosmos_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x54, 0x0a, 0x14, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc9, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x3a, 0x4c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc9, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x3a,
	0x37, 0x0a, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x3a, 0x38, 0x0a, 0x07, 0x67, 0x6f, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xcb, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x98, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x0b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cosmos_proto_cosmos_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_cosmos_proto_cosmos_proto_depIdxs = []int32{
	0, // 0: cosmos_proto.implements_interface:extendee -> google.protobuf.MessageOptions
	1, // 1: cosmos_proto.accepts_interface:extendee -> google.protobuf.FieldOptions
	1, // 2: cosmos_proto.scalar:extendee -> google.protobuf.FieldOptions
	1, // 3: cosmos_proto.go_type:extendee -> google.protobuf.FieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_proto_cosmos_proto_init() }
func file_cosmos_proto_cosmos_proto_init() {
	if File_cosmos_proto_cosmos_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_proto_cosmos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_proto_cosmos_proto_goTypes,
		DependencyIndexes: file_cosmos_proto_cosmos_proto_depIdxs,
		ExtensionInfos:    file_cosmos_proto_cosmos_proto_extTypes,
	}.Build()
	File_cosmos_proto_cosmos_proto = out.File
	file_cosmos_proto_cosmos_proto_rawDesc = nil
	file_cosmos_proto_cosmos_proto_goTypes = nil
	file_cosmos_proto_cosmos_proto_depIdxs = nil
}