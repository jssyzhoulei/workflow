package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiI
	GetPermission() permissionApiInterface

}

type apis struct {
	userApiI
	groupApiI
	permissionApiInterface permissionApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetGroupApis() groupApiI {
	return a.groupApiI
}

func (a *apis) GetPermission() permissionApiInterface {
	return a.permissionApiInterface
}
