// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: service_economic_school.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_economic_school_proto protoreflect.FileDescriptor

var file_service_economic_school_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x5f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x5f,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf1, 0x21, 0x0a, 0x0e, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x4f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x12, 0x07, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x34, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e,
	0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x32, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x2a, 0x12, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x99, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x32, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a,
	0x92, 0x41, 0x2a, 0x12, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12, 0xb7, 0x01,
	0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x62, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x32, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3c, 0x12, 0x14, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x20, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x2a, 0x12, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12, 0xa3,
	0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x4d, 0x12, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x3f, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x26, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xab, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6e, 0x65,
	0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x38, 0x12, 0x12, 0x52, 0x65, 0x6e, 0x65,
	0x77, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x22,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x72, 0x65, 0x6e, 0x65, 0x77, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x99, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x59, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a,
	0x92, 0x41, 0x3b, 0x12, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x1a, 0x2b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73,
	0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x2f,
	0x7b, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x50, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x78, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x78, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x12, 0x55, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x65,
	0x78, 0x61, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x04, 0x65, 0x78,
	0x61, 0x6d, 0x12, 0x58, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x2d, 0x75, 0x72, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0xb2, 0x01, 0x0a, 0x0e, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x77, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x65, 0x78, 0x61, 0x6d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x55, 0x12, 0x1c, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x27, 0x73, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x1a, 0x35, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x27, 0x73, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0xab, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x76, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x65, 0x78, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x45, 0x12, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x1a, 0x2d,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27,
	0x73, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0xed, 0x01,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x78, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x8b, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x92, 0x41, 0x58, 0x12, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x78, 0x61,
	0x6d, 0x73, 0x1a, 0x39, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x12, 0xae, 0x01,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x74, 0x69, 0x65, 0x73, 0x92, 0x41, 0x44, 0x12, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x2f, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x81,
	0x02, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa6, 0x01, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x74, 0x79, 0x92, 0x41, 0x7b, 0x12, 0x1c, 0x47, 0x65, 0x74, 0x20, 0x65, 0x78, 0x61,
	0x6d, 0x20, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x74, 0x79, 0x1a, 0x5b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20,
	0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x74, 0x69, 0x65, 0x73, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x74, 0x79, 0x12, 0xac, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x41,
	0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x20, 0x72, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x2b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x20, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0xcd, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x92,
	0x41, 0x4f, 0x12, 0x19, 0x47, 0x65, 0x74, 0x20, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x20, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x20, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x32, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x77, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x20, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x20, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0xc0, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x49, 0x12, 0x16, 0x47, 0x65, 0x74,
	0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x20, 0x72, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x1a, 0x2f, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x20, 0x72, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0xe2, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x83, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x92,
	0x41, 0x57, 0x12, 0x1d, 0x47, 0x65, 0x74, 0x20, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x20, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x20, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x1a, 0x36, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74,
	0x79, 0x20, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0xb1, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x76, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x92, 0x41, 0x48, 0x12, 0x19, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x2b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x20, 0x62, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x9f, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x6a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2f, 0x7b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x7d, 0x92, 0x41, 0x42, 0x12, 0x16, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x28, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20,
	0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x7d, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0xb8,
	0x01, 0x0a, 0x12, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3d, 0x12, 0x14, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x72, 0x65, 0x73, 0x65, 0x74, 0x20, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x1a, 0x25, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73,
	0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x82, 0x01, 0x5a, 0x12, 0x65, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x2d, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x62,
	0x92, 0x41, 0x6b, 0x12, 0x69, 0x0a, 0x13, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x20,
	0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x20, 0x41, 0x50, 0x49, 0x22, 0x4d, 0x0a, 0x0f, 0x59, 0x6f,
	0x73, 0x68, 0x69, 0x6b, 0x61, 0x77, 0x61, 0x20, 0x4b, 0x65, 0x69, 0x74, 0x61, 0x12, 0x22, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x77, 0x61, 0x2d, 0x4b, 0x65, 0x69, 0x74,
	0x61, 0x1a, 0x16, 0x6b, 0x65, 0x69, 0x74, 0x61, 0x5f, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x40, 0x69,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_economic_school_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                      // 0: google.protobuf.Empty
	(*CreateUserRequest)(nil),                  // 1: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),                  // 2: pb.UpdateUserRequest
	(*UpdateUserEmailRequest)(nil),             // 3: pb.UpdateUserEmailRequest
	(*VerifyChangedEmailRequest)(nil),          // 4: pb.VerifyChangedEmailRequest
	(*DeleteUserRequest)(nil),                  // 5: pb.DeleteUserRequest
	(*LoginUserRequest)(nil),                   // 6: pb.LoginUserRequest
	(*RenewAccessTokenRequest)(nil),            // 7: pb.RenewAccessTokenRequest
	(*VerifyEmailRequest)(nil),                 // 8: pb.VerifyEmailRequest
	(*CreateExamRequest)(nil),                  // 9: pb.CreateExamRequest
	(*GetExamRequest)(nil),                     // 10: pb.GetExamRequest
	(*ListExamsRequest)(nil),                   // 11: pb.ListExamsRequest
	(*UpdateExamRequest)(nil),                  // 12: pb.UpdateExamRequest
	(*DeleteExamRequest)(nil),                  // 13: pb.DeleteExamRequest
	(*GetSignedUrlRequest)(nil),                // 14: pb.GetSignedUrlRequest
	(*UpsertUserExamRequest)(nil),              // 15: pb.UpsertUserExamRequest
	(*GetUserExamRequest)(nil),                 // 16: pb.GetUserExamRequest
	(*ListCompletedUserExamsRequest)(nil),      // 17: pb.ListCompletedUserExamsRequest
	(*GetUserByUsernameParam)(nil),             // 18: pb.GetUserByUsernameParam
	(*GetUserByEmailParam)(nil),                // 19: pb.GetUserByEmailParam
	(*SendPasswordResetEmailRequest)(nil),      // 20: pb.SendPasswordResetEmailRequest
	(*PasswordResetEmailRequest)(nil),          // 21: pb.PasswordResetEmailRequest
	(*HealthCheckResponse)(nil),                // 22: pb.HealthCheckResponse
	(*CreateUserResponse)(nil),                 // 23: pb.CreateUserResponse
	(*UpdateUserResponse)(nil),                 // 24: pb.UpdateUserResponse
	(*UpdateUserEmailResponse)(nil),            // 25: pb.UpdateUserEmailResponse
	(*VerifyChangedEmailResponse)(nil),         // 26: pb.VerifyChangedEmailResponse
	(*LoginUserResponse)(nil),                  // 27: pb.LoginUserResponse
	(*RenewAccessTokenResponse)(nil),           // 28: pb.RenewAccessTokenResponse
	(*VerifyEmailResponse)(nil),                // 29: pb.VerifyEmailResponse
	(*Exam)(nil),                               // 30: pb.Exam
	(*ListExamsResponse)(nil),                  // 31: pb.ListExamsResponse
	(*GetSignedUrlResponse)(nil),               // 32: pb.GetSignedUrlResponse
	(*UserExam)(nil),                           // 33: pb.UserExam
	(*ListCompletedUserExamsResponse)(nil),     // 34: pb.ListCompletedUserExamsResponse
	(*ListUniversitiesResponse)(nil),           // 35: pb.ListUniversitiesResponse
	(*GetExamCountByUniversityResponse)(nil),   // 36: pb.GetExamCountByUniversityResponse
	(*GetGlobalRankingResponse)(nil),           // 37: pb.GetGlobalRankingResponse
	(*GetWeeklyGlobalRankingResponse)(nil),     // 38: pb.GetWeeklyGlobalRankingResponse
	(*GetUniversityRankingResponse)(nil),       // 39: pb.GetUniversityRankingResponse
	(*GetWeeklyUniversityRankingResponse)(nil), // 40: pb.GetWeeklyUniversityRankingResponse
	(*User)(nil),                               // 41: pb.User
	(*PasswordResetEmailResponse)(nil),         // 42: pb.PasswordResetEmailResponse
}
var file_service_economic_school_proto_depIdxs = []int32{
	0,  // 0: pb.EconomicSchool.HealthCheck:input_type -> google.protobuf.Empty
	1,  // 1: pb.EconomicSchool.CreateUser:input_type -> pb.CreateUserRequest
	2,  // 2: pb.EconomicSchool.UpdateUser:input_type -> pb.UpdateUserRequest
	3,  // 3: pb.EconomicSchool.UpdateUserEmail:input_type -> pb.UpdateUserEmailRequest
	4,  // 4: pb.EconomicSchool.VerifyChangedEmail:input_type -> pb.VerifyChangedEmailRequest
	5,  // 5: pb.EconomicSchool.DeleteUser:input_type -> pb.DeleteUserRequest
	6,  // 6: pb.EconomicSchool.LoginUser:input_type -> pb.LoginUserRequest
	7,  // 7: pb.EconomicSchool.RenewAccessToken:input_type -> pb.RenewAccessTokenRequest
	8,  // 8: pb.EconomicSchool.VerifyEmail:input_type -> pb.VerifyEmailRequest
	9,  // 9: pb.EconomicSchool.CreateExam:input_type -> pb.CreateExamRequest
	10, // 10: pb.EconomicSchool.GetExam:input_type -> pb.GetExamRequest
	11, // 11: pb.EconomicSchool.ListExams:input_type -> pb.ListExamsRequest
	12, // 12: pb.EconomicSchool.UpdateExam:input_type -> pb.UpdateExamRequest
	13, // 13: pb.EconomicSchool.DeleteExam:input_type -> pb.DeleteExamRequest
	14, // 14: pb.EconomicSchool.GetSignedUrl:input_type -> pb.GetSignedUrlRequest
	15, // 15: pb.EconomicSchool.UpsertUserExam:input_type -> pb.UpsertUserExamRequest
	16, // 16: pb.EconomicSchool.GetUserExam:input_type -> pb.GetUserExamRequest
	17, // 17: pb.EconomicSchool.ListCompletedUserExams:input_type -> pb.ListCompletedUserExamsRequest
	0,  // 18: pb.EconomicSchool.ListUniversities:input_type -> google.protobuf.Empty
	0,  // 19: pb.EconomicSchool.GetExamCountByUniversity:input_type -> google.protobuf.Empty
	0,  // 20: pb.EconomicSchool.GetGlobalRanking:input_type -> google.protobuf.Empty
	0,  // 21: pb.EconomicSchool.GetWeeklyGlobalRanking:input_type -> google.protobuf.Empty
	0,  // 22: pb.EconomicSchool.GetUniversityRanking:input_type -> google.protobuf.Empty
	0,  // 23: pb.EconomicSchool.GetWeeklyUniversityRanking:input_type -> google.protobuf.Empty
	18, // 24: pb.EconomicSchool.GetUserByUsername:input_type -> pb.GetUserByUsernameParam
	19, // 25: pb.EconomicSchool.GetUserByEmail:input_type -> pb.GetUserByEmailParam
	20, // 26: pb.EconomicSchool.SendPasswordResetEmail:input_type -> pb.SendPasswordResetEmailRequest
	21, // 27: pb.EconomicSchool.PasswordResetEmail:input_type -> pb.PasswordResetEmailRequest
	22, // 28: pb.EconomicSchool.HealthCheck:output_type -> pb.HealthCheckResponse
	23, // 29: pb.EconomicSchool.CreateUser:output_type -> pb.CreateUserResponse
	24, // 30: pb.EconomicSchool.UpdateUser:output_type -> pb.UpdateUserResponse
	25, // 31: pb.EconomicSchool.UpdateUserEmail:output_type -> pb.UpdateUserEmailResponse
	26, // 32: pb.EconomicSchool.VerifyChangedEmail:output_type -> pb.VerifyChangedEmailResponse
	0,  // 33: pb.EconomicSchool.DeleteUser:output_type -> google.protobuf.Empty
	27, // 34: pb.EconomicSchool.LoginUser:output_type -> pb.LoginUserResponse
	28, // 35: pb.EconomicSchool.RenewAccessToken:output_type -> pb.RenewAccessTokenResponse
	29, // 36: pb.EconomicSchool.VerifyEmail:output_type -> pb.VerifyEmailResponse
	30, // 37: pb.EconomicSchool.CreateExam:output_type -> pb.Exam
	30, // 38: pb.EconomicSchool.GetExam:output_type -> pb.Exam
	31, // 39: pb.EconomicSchool.ListExams:output_type -> pb.ListExamsResponse
	30, // 40: pb.EconomicSchool.UpdateExam:output_type -> pb.Exam
	0,  // 41: pb.EconomicSchool.DeleteExam:output_type -> google.protobuf.Empty
	32, // 42: pb.EconomicSchool.GetSignedUrl:output_type -> pb.GetSignedUrlResponse
	33, // 43: pb.EconomicSchool.UpsertUserExam:output_type -> pb.UserExam
	33, // 44: pb.EconomicSchool.GetUserExam:output_type -> pb.UserExam
	34, // 45: pb.EconomicSchool.ListCompletedUserExams:output_type -> pb.ListCompletedUserExamsResponse
	35, // 46: pb.EconomicSchool.ListUniversities:output_type -> pb.ListUniversitiesResponse
	36, // 47: pb.EconomicSchool.GetExamCountByUniversity:output_type -> pb.GetExamCountByUniversityResponse
	37, // 48: pb.EconomicSchool.GetGlobalRanking:output_type -> pb.GetGlobalRankingResponse
	38, // 49: pb.EconomicSchool.GetWeeklyGlobalRanking:output_type -> pb.GetWeeklyGlobalRankingResponse
	39, // 50: pb.EconomicSchool.GetUniversityRanking:output_type -> pb.GetUniversityRankingResponse
	40, // 51: pb.EconomicSchool.GetWeeklyUniversityRanking:output_type -> pb.GetWeeklyUniversityRankingResponse
	41, // 52: pb.EconomicSchool.GetUserByUsername:output_type -> pb.User
	41, // 53: pb.EconomicSchool.GetUserByEmail:output_type -> pb.User
	0,  // 54: pb.EconomicSchool.SendPasswordResetEmail:output_type -> google.protobuf.Empty
	42, // 55: pb.EconomicSchool.PasswordResetEmail:output_type -> pb.PasswordResetEmailResponse
	28, // [28:56] is the sub-list for method output_type
	0,  // [0:28] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_economic_school_proto_init() }
func file_service_economic_school_proto_init() {
	if File_service_economic_school_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_login_user_proto_init()
	file_rpc_update_user_proto_init()
	file_rpc_verify_email_proto_init()
	file_message_economic_school_proto_init()
	file_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_economic_school_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_economic_school_proto_goTypes,
		DependencyIndexes: file_service_economic_school_proto_depIdxs,
	}.Build()
	File_service_economic_school_proto = out.File
	file_service_economic_school_proto_rawDesc = nil
	file_service_economic_school_proto_goTypes = nil
	file_service_economic_school_proto_depIdxs = nil
}
