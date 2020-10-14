package client

import (
	"gitee.com/grandeep/org-svc/client/parser"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func addUserGrpcConn(conn *grpc.ClientConn) services.UserServiceInterface {
	return &endpoints.UserServiceEndpoint{
		AddUserEndpoint:     grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddUser",
			parser.EncodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.NullResponse{},
		).Endpoint(),
	}
}

func groupAddGrpcConn(conn *grpc.ClientConn) services.GroupServiceI {
	return &endpoints.GroupServiceEndpoint{
		GroupAddEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupAdd",
			parser.EncodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.GroupResponse{},
			).Endpoint(),
	}
}
