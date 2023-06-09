// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: auth/delete_auth.proto

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
	DeleteAuthService_DeleteAuth_FullMethodName = "/DeleteAuthService/DeleteAuth"
)

// DeleteAuthServiceClient is the client API for DeleteAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteAuthServiceClient interface {
	DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error)
}

type deleteAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteAuthServiceClient(cc grpc.ClientConnInterface) DeleteAuthServiceClient {
	return &deleteAuthServiceClient{cc}
}

func (c *deleteAuthServiceClient) DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error) {
	out := new(DeleteAuthResponse)
	err := c.cc.Invoke(ctx, DeleteAuthService_DeleteAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteAuthServiceServer is the server API for DeleteAuthService service.
// All implementations should embed UnimplementedDeleteAuthServiceServer
// for forward compatibility
type DeleteAuthServiceServer interface {
	DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error)
}

// UnimplementedDeleteAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDeleteAuthServiceServer struct {
}

func (UnimplementedDeleteAuthServiceServer) DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuth not implemented")
}

// UnsafeDeleteAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteAuthServiceServer will
// result in compilation errors.
type UnsafeDeleteAuthServiceServer interface {
	mustEmbedUnimplementedDeleteAuthServiceServer()
}

func RegisterDeleteAuthServiceServer(s grpc.ServiceRegistrar, srv DeleteAuthServiceServer) {
	s.RegisterService(&DeleteAuthService_ServiceDesc, srv)
}

func _DeleteAuthService_DeleteAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteAuthServiceServer).DeleteAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeleteAuthService_DeleteAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteAuthServiceServer).DeleteAuth(ctx, req.(*DeleteAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteAuthService_ServiceDesc is the grpc.ServiceDesc for DeleteAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeleteAuthService",
	HandlerType: (*DeleteAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteAuth",
			Handler:    _DeleteAuthService_DeleteAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/delete_auth.proto",
}
