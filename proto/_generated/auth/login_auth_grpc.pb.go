// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: auth/login_auth.proto

package auth

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
	LoginAuthService_LoginAuth_FullMethodName = "/LoginAuthService/LoginAuth"
)

// LoginAuthServiceClient is the client API for LoginAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginAuthServiceClient interface {
	LoginAuth(ctx context.Context, in *LoginAuthRequest, opts ...grpc.CallOption) (*LoginAuthResponse, error)
}

type loginAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginAuthServiceClient(cc grpc.ClientConnInterface) LoginAuthServiceClient {
	return &loginAuthServiceClient{cc}
}

func (c *loginAuthServiceClient) LoginAuth(ctx context.Context, in *LoginAuthRequest, opts ...grpc.CallOption) (*LoginAuthResponse, error) {
	out := new(LoginAuthResponse)
	err := c.cc.Invoke(ctx, LoginAuthService_LoginAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginAuthServiceServer is the server API for LoginAuthService service.
// All implementations should embed UnimplementedLoginAuthServiceServer
// for forward compatibility
type LoginAuthServiceServer interface {
	LoginAuth(context.Context, *LoginAuthRequest) (*LoginAuthResponse, error)
}

// UnimplementedLoginAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLoginAuthServiceServer struct {
}

func (UnimplementedLoginAuthServiceServer) LoginAuth(context.Context, *LoginAuthRequest) (*LoginAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAuth not implemented")
}

// UnsafeLoginAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginAuthServiceServer will
// result in compilation errors.
type UnsafeLoginAuthServiceServer interface {
	mustEmbedUnimplementedLoginAuthServiceServer()
}

func RegisterLoginAuthServiceServer(s grpc.ServiceRegistrar, srv LoginAuthServiceServer) {
	s.RegisterService(&LoginAuthService_ServiceDesc, srv)
}

func _LoginAuthService_LoginAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginAuthServiceServer).LoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginAuthService_LoginAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginAuthServiceServer).LoginAuth(ctx, req.(*LoginAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginAuthService_ServiceDesc is the grpc.ServiceDesc for LoginAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LoginAuthService",
	HandlerType: (*LoginAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginAuth",
			Handler:    _LoginAuthService_LoginAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/login_auth.proto",
}
