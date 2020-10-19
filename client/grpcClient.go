package client

import (
	"gitee.com/grandeep/org-svc/client/parser"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func userGrpcConn(conn *grpc.ClientConn) services.UserServiceInterface {
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
			"RpcGetUserById",
			parser.EncodeIndexProto,
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
		AddUsersEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddUsers",
			parser.EncodeAddUsersRequest,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
	}
}

// groupGrpcConn 组
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
		GroupUpdateEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupUpdate",
			parser.EncodeGroupUpdateProto,
			parser.DecodeGroupProto,
			pb_user_v1.GroupResponse{},
		).Endpoint(),
		QuotaUpdateEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQuotaUpdate",
			parser.EncodeQuotaUpdateProto,
			parser.DecodeGroupProto,
			pb_user_v1.GroupResponse{},
		).Endpoint(),
	}
}

func RoleGrpcConn(conn *grpc.ClientConn) services.RoleServiceI {
	return &endpoints.RoleServiceEndpoint{
		AddRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddRole",
			parser.EncodeCreateMenuPermRequestModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		UpdateRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcUpdateRole",
			parser.EncodeCreateMenuPermRequestModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		DeleteRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcDeleteRole",
			parser.EncodeCreateMenuPermRequestModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
	}
}

func permissionGrpcConn(conn *grpc.ClientConn) services.PermissionServiceInterface {
	return &endpoints.PermissionServiceEndpoint{
		AddPermissionEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddPermission",
			parser.EncodePermissionModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		AddMenuEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddMenu",
			parser.EncodeMenuModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		GetMenuCascadeByModuleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcGetMenuCascadeByModule",
			parser.EncodeMenuModule,
			parser.DecodeCascadeProto,
			pb_user_v1.Cascades{},
		).Endpoint(),
		GetPermissionByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcGetPermissionByID",
			parser.EncodeIndexProto,
			parser.DecodePermissionProto,
			pb_user_v1.PermissionProto{},
		).Endpoint(),
		DeletePermissionByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcDeletePermissionByID",
			parser.EncodeIndexProto,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		UpdatePermissionByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcUpdatePermissionByID",
			parser.EncodePermissionModel,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
	}
}

