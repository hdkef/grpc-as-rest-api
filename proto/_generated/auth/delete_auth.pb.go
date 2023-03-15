// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: auth/delete_auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "grpcrest/proto/_generated/google/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteAuthRequest) Reset() {
	*x = DeleteAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_delete_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthRequest) ProtoMessage() {}

func (x *DeleteAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_delete_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthRequest) Descriptor() ([]byte, []int) {
	return file_auth_delete_auth_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteAuthRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteAuthResponse) Reset() {
	*x = DeleteAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_delete_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthResponse) ProtoMessage() {}

func (x *DeleteAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_delete_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_delete_auth_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteAuthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_auth_delete_auth_proto protoreflect.FileDescriptor

var file_auth_delete_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x4a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x33, 0x42, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x72, 0x65, 0x73, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_delete_auth_proto_rawDescOnce sync.Once
	file_auth_delete_auth_proto_rawDescData = file_auth_delete_auth_proto_rawDesc
)

func file_auth_delete_auth_proto_rawDescGZIP() []byte {
	file_auth_delete_auth_proto_rawDescOnce.Do(func() {
		file_auth_delete_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_delete_auth_proto_rawDescData)
	})
	return file_auth_delete_auth_proto_rawDescData
}

var file_auth_delete_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_delete_auth_proto_goTypes = []interface{}{
	(*DeleteAuthRequest)(nil),  // 0: DeleteAuthRequest
	(*DeleteAuthResponse)(nil), // 1: DeleteAuthResponse
}
var file_auth_delete_auth_proto_depIdxs = []int32{
	0, // 0: DeleteAuthService.DeleteAuth:input_type -> DeleteAuthRequest
	1, // 1: DeleteAuthService.DeleteAuth:output_type -> DeleteAuthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_delete_auth_proto_init() }
func file_auth_delete_auth_proto_init() {
	if File_auth_delete_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_delete_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthRequest); i {
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
		file_auth_delete_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthResponse); i {
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
			RawDescriptor: file_auth_delete_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_delete_auth_proto_goTypes,
		DependencyIndexes: file_auth_delete_auth_proto_depIdxs,
		MessageInfos:      file_auth_delete_auth_proto_msgTypes,
	}.Build()
	File_auth_delete_auth_proto = out.File
	file_auth_delete_auth_proto_rawDesc = nil
	file_auth_delete_auth_proto_goTypes = nil
	file_auth_delete_auth_proto_depIdxs = nil
}
