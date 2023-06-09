// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user/get_all_user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GetAllUserService_GetAllUser_FullMethodName = "/GetAllUserService/GetAllUser"
)

// GetAllUserServiceClient is the client API for GetAllUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetAllUserServiceClient interface {
	GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserResponse, error)
}

type getAllUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetAllUserServiceClient(cc grpc.ClientConnInterface) GetAllUserServiceClient {
	return &getAllUserServiceClient{cc}
}

func (c *getAllUserServiceClient) GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserResponse, error) {
	out := new(GetAllUserResponse)
	err := c.cc.Invoke(ctx, GetAllUserService_GetAllUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetAllUserServiceServer is the server API for GetAllUserService service.
// All implementations should embed UnimplementedGetAllUserServiceServer
// for forward compatibility
type GetAllUserServiceServer interface {
	GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserResponse, error)
}

// UnimplementedGetAllUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGetAllUserServiceServer struct {
}

func (UnimplementedGetAllUserServiceServer) GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}

// UnsafeGetAllUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetAllUserServiceServer will
// result in compilation errors.
type UnsafeGetAllUserServiceServer interface {
	mustEmbedUnimplementedGetAllUserServiceServer()
}

func RegisterGetAllUserServiceServer(s grpc.ServiceRegistrar, srv GetAllUserServiceServer) {
	s.RegisterService(&GetAllUserService_ServiceDesc, srv)
}

func _GetAllUserService_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetAllUserServiceServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetAllUserService_GetAllUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetAllUserServiceServer).GetAllUser(ctx, req.(*GetAllUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetAllUserService_ServiceDesc is the grpc.ServiceDesc for GetAllUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetAllUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetAllUserService",
	HandlerType: (*GetAllUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUser",
			Handler:    _GetAllUserService_GetAllUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/get_all_user.proto",
}
