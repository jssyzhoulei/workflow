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
			parser.EncodeUserRoleDTO,
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
			parser.EncodeUserRoleDTO,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		DeleteUserByIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcDeleteUserByID",
			parser.EncodeIndexProto,
			parser.DecodeNullProto,
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
		GetUserListEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcGetUserList",
			parser.EncodeUserPage,
			parser.DecodeUsersPage,
			pb_user_v1.UsersPage{},
		).Endpoint(),
		BatchDeleteUsersEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcBatchDeleteUsers",
			parser.EncodeUserIDs,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		ImportUsersByGroupIdEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcImportUsersByGroupId",
			parser.EncodeGroupAndUserId,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		GetUsersEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcGetUsers",
			parser.EncodeUserCondition,
			parser.DecodeUserResponse,
			pb_user_v1.UserQueryResponse{},
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
		GroupTreeQueryEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupTreeQuery",
			parser.EncodeGroupIDProto,
			parser.DecodeGroupTreeQueryProto,
			pb_user_v1.GroupTreeResponse{},
		).Endpoint(),
		GroupDeleteEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCGroupDelete",
			parser.EncodeGroupIDProto,
			parser.DecodeGroupProto,
			pb_user_v1.GroupResponse{},
		).Endpoint(),
		QueryGroupAndSubGroupsUsersEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQueryGroupAndSubGroupsUsers",
			parser.EncodeGroupIDProto,
			parser.DecodeUsers,
			pb_user_v1.Users{},
		).Endpoint(),
		SetGroupQuotaUsedEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCSetGroupQuotaUsed",
			parser.EncodeSetGroupQuotaUsedProto,
			parser.DecodeGroupProto,
			pb_user_v1.GroupResponse{},
		).Endpoint(),
		QueryGroupIDAndSubGroupsIDEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQueryGroupIDAndSubGroupsID",
			parser.EncodeGroupIDProto,
			parser.DecodeGroupIDsResponse,
			pb_user_v1.GroupIDsResponse{},
		).Endpoint(),
		QuerySubGroupsUsersEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQuerySubGroupsUsers",
			parser.EncodeGroupIDProto,
			parser.DecodeUsers,
			pb_user_v1.Users{},
		).Endpoint(),
		QueryQuotaByConditionEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQueryQuotaByCondition",
			parser.EncodeQueryQuotaByCondition,
			parser.DecodeQueryQuotaByConditionResponse,
			pb_user_v1.QueryQuotaByConditionResponse{},
		).Endpoint(),
		GetAllGroupsEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcGetGroups",
			parser.EncodeGroupIDProto,
			parser.DecodeGroupsProto,
			pb_user_v1.Groups{},
		).Endpoint(),
		QueryQuotaEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQueryQuota",
			parser.EncodeGroupIDProto,
			parser.DecodeQueryQuotaResponse,
			pb_user_v1.QueryQuotaResponse{},
		).Endpoint(),
		QueryTopGroupExcludeSelfUsersEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RPCQueryTopGroupExcludeSelfUsers",
			parser.EncodeGroupIDWithPage,
			parser.DecodeGroupUsersWithPage,
			pb_user_v1.GroupUsersWithPage{},
		).Endpoint(),
	}
}

func RoleGrpcConn(conn *grpc.ClientConn) services.RoleServiceI {
	return &endpoints.RoleServiceEndpoint{
		AddRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcAddRole",
			parser.EncodeCreateMenuPermRequest,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		UpdateRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcUpdateRole",
			parser.EncodeCreateMenuPermRequest,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		DeleteRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcDeleteRole",
			parser.EncodeIndexProto,
			parser.DecodeNullProto,
			pb_user_v1.NullResponse{},
		).Endpoint(),
		QueryRoleEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcQueryRole",
			parser.EncodeIndexProto,
			parser.DecodeCreateMenuPermRequestProto,
			pb_user_v1.CreateMenuPermRequestProto{},
		).Endpoint(),
		QueryRolesEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcQueryRoles",
			parser.EncodeRolePageProto,
			parser.DecodeRolePageProto,
			pb_user_v1.RolePageRequestProto{},
		).Endpoint(),
		MenuTreeEndpoint: grpctransport.NewClient(
			conn,
			"pb_user_v1.RpcOrgService",
			"RpcMenuTree",
			parser.EncodeMenuModule,
			parser.DecodeCascadeProto,
			pb_user_v1.Cascades{},
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
