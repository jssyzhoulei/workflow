package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiInterface
	GetPermission() permissionApiInterface
	GetGroupApis() groupAPIInterface
	GetRoleApis() RoleApiInterface
}

type apis struct {
	userApiInterface
	permissionApiInterface permissionApiInterface
	groupAPIInterface
	RoleApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiInterface: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
		//groupApiI,NewGroupApi(o.),
		RoleApiInterface:NewRoleApi(o.GetRoleService()),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
		groupAPIInterface: NewGroupAPI(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiInterface {
	return a.userApiInterface
}

func (a *apis) GetGroupApis() groupAPIInterface {
	return a.groupAPIInterface
}
func (a *apis) GetRoleApis() RoleApiInterface {
	return a.RoleApiInterface
}

func (a *apis) GetPermission() permissionApiInterface {
	return a.permissionApiInterface
}