package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiI
	GetGroupApis() groupApiInterface

}

type apis struct {
	userApiI
	groupApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		groupApiInterface: NewGroupApi(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetGroupApis() groupApiInterface {
	return a.groupApiInterface
}
