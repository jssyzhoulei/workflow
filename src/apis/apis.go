package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiI
	GetPermission() permissionApiInterface
	GetRoleApis() RoleApiInterface
	GetGroupApis() groupApiInterface
}

type apis struct {
	userApiI
	permissionApiInterface permissionApiInterface
	RoleApiInterface
	groupApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		RoleApiInterface:NewRoleApi(o.GetRoleService()),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
		groupApiInterface: NewGroupApi(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetRoleApis() RoleApiInterface {
	return a.RoleApiInterface
}

func (a *apis) GetGroupApis() groupApiInterface {
	return a.groupApiInterface
}

func (a *apis) GetPermission() permissionApiInterface {
	return a.permissionApiInterface
}
