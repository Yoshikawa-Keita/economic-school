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
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf3, 0x07, 0x0a, 0x0e, 0x45, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x8e, 0x01, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x34, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12, 0x84, 0x01,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x32, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x2a, 0x12, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x12, 0xa3, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x69, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x4d, 0x12, 0x0a, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x3f, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x26, 0x20, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x99, 0x01, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3b, 0x12, 0x0c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2b, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x44, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x4b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x73, 0x12, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73, 0x12, 0x55,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x73,
	0x2f, 0x7b, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x3a,
	0x04, 0x65, 0x78, 0x61, 0x6d, 0x12, 0x58, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x78, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x82, 0x01, 0x5a, 0x12, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x2d, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x6b, 0x12, 0x69, 0x0a, 0x13, 0x45, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x20, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x20, 0x41, 0x50, 0x49,
	0x22, 0x4d, 0x0a, 0x0f, 0x59, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x77, 0x61, 0x20, 0x4b, 0x65,
	0x69, 0x74, 0x61, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x77,
	0x61, 0x2d, 0x4b, 0x65, 0x69, 0x74, 0x61, 0x1a, 0x16, 0x6b, 0x65, 0x69, 0x74, 0x61, 0x5f, 0x79,
	0x6f, 0x73, 0x68, 0x69, 0x40, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x32,
	0x03, 0x31, 0x2e, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_economic_school_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),   // 0: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),   // 1: pb.UpdateUserRequest
	(*LoginUserRequest)(nil),    // 2: pb.LoginUserRequest
	(*VerifyEmailRequest)(nil),  // 3: pb.VerifyEmailRequest
	(*CreateExamRequest)(nil),   // 4: pb.CreateExamRequest
	(*GetExamRequest)(nil),      // 5: pb.GetExamRequest
	(*ListExamsRequest)(nil),    // 6: pb.ListExamsRequest
	(*UpdateExamRequest)(nil),   // 7: pb.UpdateExamRequest
	(*DeleteExamRequest)(nil),   // 8: pb.DeleteExamRequest
	(*CreateUserResponse)(nil),  // 9: pb.CreateUserResponse
	(*UpdateUserResponse)(nil),  // 10: pb.UpdateUserResponse
	(*LoginUserResponse)(nil),   // 11: pb.LoginUserResponse
	(*VerifyEmailResponse)(nil), // 12: pb.VerifyEmailResponse
	(*Exam)(nil),                // 13: pb.Exam
	(*ListExamsResponse)(nil),   // 14: pb.ListExamsResponse
	(*emptypb.Empty)(nil),       // 15: google.protobuf.Empty
}
var file_service_economic_school_proto_depIdxs = []int32{
	0,  // 0: pb.EconomicSchool.CreateUser:input_type -> pb.CreateUserRequest
	1,  // 1: pb.EconomicSchool.UpdateUser:input_type -> pb.UpdateUserRequest
	2,  // 2: pb.EconomicSchool.LoginUser:input_type -> pb.LoginUserRequest
	3,  // 3: pb.EconomicSchool.VerifyEmail:input_type -> pb.VerifyEmailRequest
	4,  // 4: pb.EconomicSchool.CreateExam:input_type -> pb.CreateExamRequest
	5,  // 5: pb.EconomicSchool.GetExam:input_type -> pb.GetExamRequest
	6,  // 6: pb.EconomicSchool.ListExams:input_type -> pb.ListExamsRequest
	7,  // 7: pb.EconomicSchool.UpdateExam:input_type -> pb.UpdateExamRequest
	8,  // 8: pb.EconomicSchool.DeleteExam:input_type -> pb.DeleteExamRequest
	9,  // 9: pb.EconomicSchool.CreateUser:output_type -> pb.CreateUserResponse
	10, // 10: pb.EconomicSchool.UpdateUser:output_type -> pb.UpdateUserResponse
	11, // 11: pb.EconomicSchool.LoginUser:output_type -> pb.LoginUserResponse
	12, // 12: pb.EconomicSchool.VerifyEmail:output_type -> pb.VerifyEmailResponse
	13, // 13: pb.EconomicSchool.CreateExam:output_type -> pb.Exam
	13, // 14: pb.EconomicSchool.GetExam:output_type -> pb.Exam
	14, // 15: pb.EconomicSchool.ListExams:output_type -> pb.ListExamsResponse
	13, // 16: pb.EconomicSchool.UpdateExam:output_type -> pb.Exam
	15, // 17: pb.EconomicSchool.DeleteExam:output_type -> google.protobuf.Empty
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
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
