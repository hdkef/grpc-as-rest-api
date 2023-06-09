// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user/create_user.proto

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
	CreateUserService_CreateUser_FullMethodName = "/CreateUserService/CreateUser"
)

// CreateUserServiceClient is the client API for CreateUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateUserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type createUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateUserServiceClient(cc grpc.ClientConnInterface) CreateUserServiceClient {
	return &createUserServiceClient{cc}
}

func (c *createUserServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, CreateUserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateUserServiceServer is the server API for CreateUserService service.
// All implementations should embed UnimplementedCreateUserServiceServer
// for forward compatibility
type CreateUserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedCreateUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCreateUserServiceServer struct {
}

func (UnimplementedCreateUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// UnsafeCreateUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateUserServiceServer will
// result in compilation errors.
type UnsafeCreateUserServiceServer interface {
	mustEmbedUnimplementedCreateUserServiceServer()
}

func RegisterCreateUserServiceServer(s grpc.ServiceRegistrar, srv CreateUserServiceServer) {
	s.RegisterService(&CreateUserService_ServiceDesc, srv)
}

func _CreateUserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateUserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateUserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateUserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateUserService_ServiceDesc is the grpc.ServiceDesc for CreateUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CreateUserService",
	HandlerType: (*CreateUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _CreateUserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/create_user.proto",
}
