package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiI
	GetPermission() permissionApiInterface
	GetGroupApis() groupApiInterface

}

type apis struct {
	userApiI
	permissionApiInterface permissionApiInterface
	groupApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
		groupApiInterface: NewGroupApi(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetGroupApis() groupApiInterface {
	return a.groupApiInterface
}

func (a *apis) GetPermission() permissionApiInterface {
	return a.permissionApiInterface
}
