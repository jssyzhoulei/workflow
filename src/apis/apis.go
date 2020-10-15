package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiInterface
	GetPermission() permissionApiInterface
	GetGroupApis() groupAPIInterface

}

type apis struct {
	userApiInterface
	permissionApiInterface permissionApiInterface
	groupAPIInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiInterface: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
		//groupApiI,NewGroupApi(o.),
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

func (a *apis) GetPermission() permissionApiInterface {
	return a.permissionApiInterface
}