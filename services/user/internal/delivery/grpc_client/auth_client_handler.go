package grpcclient

import (
	authpb "grpcrest/proto/_generated/auth"

	"google.golang.org/grpc"
)

type AuthGRPCClient struct {
	Create authpb.CreateAuthServiceClient
	Delete authpb.DeleteAuthServiceClient
	Update authpb.UpdateAuthServiceClient
}

func AuthGRPCClientHandler(conn *grpc.ClientConn) (*AuthGRPCClient, error) {
	var c AuthGRPCClient
	//bind new client grpc to struct
	createClient := authpb.NewCreateAuthServiceClient(conn)
	updateClient := authpb.NewUpdateAuthServiceClient(conn)
	deleteClient := authpb.NewDeleteAuthServiceClient(conn)
	c.Create = createClient
	c.Update = updateClient
	c.Delete = deleteClient
	return &c, nil
}
