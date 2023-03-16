package grpcclient

import (
	authpb "grpcrest/proto/_generated/auth"

	"google.golang.org/grpc"
)

type AuthGRPCClient struct {
	authpb.CreateAuthServiceClient
	authpb.DeleteAuthServiceClient
	authpb.UpdateAuthServiceClient
}

type AuthGRPCClient_ interface {
	authpb.CreateAuthServiceClient
	authpb.DeleteAuthServiceClient
	authpb.UpdateAuthServiceClient
}

func NewAuthGRPCClientHandler(conn *grpc.ClientConn) (AuthGRPCClient_, error) {
	var c AuthGRPCClient
	//bind new client grpc to struct
	createClient := authpb.NewCreateAuthServiceClient(conn)
	updateClient := authpb.NewUpdateAuthServiceClient(conn)
	deleteClient := authpb.NewDeleteAuthServiceClient(conn)
	c.CreateAuthServiceClient = createClient
	c.UpdateAuthServiceClient = updateClient
	c.DeleteAuthServiceClient = deleteClient
	return &c, nil
}
