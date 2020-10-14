package client

import (
	"gitee.com/grandeep/org-svc/client/parser"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func addUserGrpcConn(conn *grpc.ClientConn) services.UserServiceI {
	return &endpoints.UserServiceEndpoint{
		AddUserEndpoint:     grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddUser",
			parser.EncodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.UserProto{},
		).Endpoint(),
	}
}

// groupAddGrpcConn 添加组
func groupAddGrpcConn(conn *grpc.ClientConn) services.GroupServiceInterface {
	return &endpoints.GroupServiceEndpoint{
		GroupAddEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupAdd",
			parser.EncodeGroupAddProto,
			parser.DecodeGroupProto,
			pb_user_v1.GroupResponse{},
			).Endpoint(),
	}
}

// groupQueryByConditionGrpcConn 按条件查询组
func groupQueryByConditionGrpcConn(conn *grpc.ClientConn) services.GroupServiceInterface {
	return &endpoints.GroupServiceEndpoint{
		GroupAddEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupQueryByCondition",
			parser.EncodeGroupQueryByConditionProto,
			parser.DecodeGroupQueryByConditionProto,
			pb_user_v1.GroupQueryByConditionResponse{},
		).Endpoint(),
	}
}

