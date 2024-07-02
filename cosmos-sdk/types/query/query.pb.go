// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/query/v1/query.proto

package query

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

var E_ModuleQuerySafe = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         11110001,
	Name:          "cosmos.query.v1.module_query_safe",
	Tag:           "varint,11110001,opt,name=module_query_safe",
	Filename:      "cosmos/query/v1/query.proto",
}

func init() {
	// proto.RegisterExtension(E_ModuleQuerySafe)
}

//func init() { proto.RegisterFile("cosmos/query/v1/query.proto", fileDescriptor_5c815d91553f8dca) }

var fileDescriptor_5c815d91553f8dca = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2f, 0x33, 0x84, 0x30, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x21, 0x92, 0x7a, 0x10, 0xb1, 0x32, 0x43, 0x29, 0x85, 0xf4, 0xfc,
	0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0xb0, 0x74, 0x52, 0x69, 0x9a, 0x7e, 0x4a, 0x6a, 0x71, 0x72, 0x51,
	0x66, 0x41, 0x49, 0x7e, 0x11, 0x44, 0x8b, 0x95, 0x2f, 0x97, 0x60, 0x6e, 0x7e, 0x4a, 0x69, 0x4e,
	0x6a, 0x3c, 0x58, 0x53, 0x7c, 0x71, 0x62, 0x5a, 0xaa, 0x90, 0x9c, 0x1e, 0x44, 0x9f, 0x1e, 0x4c,
	0x9f, 0x9e, 0x6f, 0x6a, 0x49, 0x46, 0x7e, 0x8a, 0x7f, 0x41, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0xc4,
	0xc7, 0x9e, 0x65, 0xac, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0xfc, 0x10, 0xbd, 0x81, 0x20, 0xad, 0xc1,
	0x89, 0x69, 0xa9, 0x4e, 0xde, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65,
	0x98, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x96, 0x5f, 0x94, 0x94,
	0x9f, 0x93, 0xaa, 0x9f, 0x55, 0x9a, 0x97, 0xaf, 0x5f, 0x66, 0xaa, 0x0f, 0x71, 0xb6, 0x6e, 0x71,
	0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41, 0x2a, 0xd4, 0x77, 0x49, 0x6c, 0x60, 0xeb, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x59, 0x5f, 0xf9, 0x53, 0xf4, 0x00, 0x00, 0x00,
}
