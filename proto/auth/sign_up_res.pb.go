// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.0
// source: proto/auth/sign_up_res.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignUpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SignUpRes) Reset() {
	*x = SignUpRes{}
	mi := &file_proto_auth_sign_up_res_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRes) ProtoMessage() {}

func (x *SignUpRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_sign_up_res_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRes.ProtoReflect.Descriptor instead.
func (*SignUpRes) Descriptor() ([]byte, []int) {
	return file_proto_auth_sign_up_res_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRes) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_proto_auth_sign_up_res_proto protoreflect.FileDescriptor

var file_proto_auth_sign_up_res_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6d, 0x65,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_sign_up_res_proto_rawDescOnce sync.Once
	file_proto_auth_sign_up_res_proto_rawDescData = file_proto_auth_sign_up_res_proto_rawDesc
)

func file_proto_auth_sign_up_res_proto_rawDescGZIP() []byte {
	file_proto_auth_sign_up_res_proto_rawDescOnce.Do(func() {
		file_proto_auth_sign_up_res_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_sign_up_res_proto_rawDescData)
	})
	return file_proto_auth_sign_up_res_proto_rawDescData
}

var file_proto_auth_sign_up_res_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_auth_sign_up_res_proto_goTypes = []any{
	(*SignUpRes)(nil),             // 0: proto.SignUpRes
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_proto_auth_sign_up_res_proto_depIdxs = []int32{
	1, // 0: proto.SignUpRes.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_auth_sign_up_res_proto_init() }
func file_proto_auth_sign_up_res_proto_init() {
	if File_proto_auth_sign_up_res_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auth_sign_up_res_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_auth_sign_up_res_proto_goTypes,
		DependencyIndexes: file_proto_auth_sign_up_res_proto_depIdxs,
		MessageInfos:      file_proto_auth_sign_up_res_proto_msgTypes,
	}.Build()
	File_proto_auth_sign_up_res_proto = out.File
	file_proto_auth_sign_up_res_proto_rawDesc = nil
	file_proto_auth_sign_up_res_proto_goTypes = nil
	file_proto_auth_sign_up_res_proto_depIdxs = nil
}
