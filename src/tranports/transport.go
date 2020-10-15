package tranports

import (
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

// OrgTransport 实现了 pb_user_v1.RpcOrgServiceServer
type OrgTransport struct {
	*userGrpcTransport
	*groupGrpcTransport
	*permissionGrpcTransport
	*roleGrpcTransport
}

func NewTransport(ept *endpoints.OrgEndpoint) pb_user_v1.RpcOrgServiceServer {
	return &OrgTransport{
		userGrpcTransport: NewUserGrpcTransport(ept.UserServiceEndpoint),
		groupGrpcTransport: NewGroupGrpcTransport(ept.GroupServiceEndpoint),
		permissionGrpcTransport: NewPermissionGrpcTransport(ept.PermissionServiceEndpoint),
		roleGrpcTransport: NewRoleGrpcTransport(ept.RoleServiceEndpoint),
	}
}
