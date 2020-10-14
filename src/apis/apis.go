package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiInterface

}

type apis struct {
	userApiInterface
	groupApiI
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiInterface: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
	}
}

func (a *apis) GetUserApis() userApiInterface {
	return a.userApiInterface
}

func (a *apis) GetGroupApis() groupApiI {
	return a.groupApiI
}
