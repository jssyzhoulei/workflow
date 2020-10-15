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
		AddUserEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddUser",
			parser.EncodeUserModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		GetUserByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcGetUserByID",
			parser.EncodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.UserProto{},
			).Endpoint(),
		UpdateUserByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcUpdateUserByID",
			parser.EncodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.NullResponse{},
			).Endpoint(),
		DeleteUserByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcDeleteUserByID",
			parser.DecodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.NullResponse{},
			).Endpoint(),
	}
}

// groupAddGrpcConn 添加组
func groupGrpcConn(conn *grpc.ClientConn) services.GroupServiceInterface {
	return &endpoints.GroupServiceEndpoint{
		GroupAddEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupAdd",
			parser.EncodeGroupAddProto,
			parser.DecodeGroupProto,
			pb_user_v1.GroupResponse{},
		).Endpoint(),
		GroupQueryWithQuotaByConditionEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupQueryWithQuotaByCondition",
			parser.EncodeGroupQueryWithQuotaByConditionProto,
			parser.DecodeGroupQueryWithQuotaByConditionProto,
			pb_user_v1.GroupQueryWithQuotaByConditionResponse{},
		).Endpoint(),
	}
}

func addRoleGrpcConn(conn *grpc.ClientConn) services.RoleServiceI {
	return &endpoints.RoleServiceEndpoint{
		AddRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddRole",
			parser.DecodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.RoleProto{},
		).Endpoint(),
		UpdateRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcUpdateRole",
			parser.DecodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.RoleProto{},
		).Endpoint(),
		DeleteRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcDeleteRole",
			parser.DecodeUserModel,
			parser.DecodeUserModel,
			pb_user_v1.RoleProto{},
		).Endpoint(),
	}
}

func permissionGrpcConn(conn *grpc.ClientConn) services.PermissionServiceInterface {
	return &endpoints.PermissionServiceEndpoint{
		AddPermissionEndpoint:          grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddPermission",
			parser.EncodePermissionModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
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
