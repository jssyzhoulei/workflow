package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiI
	GetGroupApis() groupAPIInterface

}

type apis struct {
	userApiI
	groupAPIInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		groupAPIInterface: NewGroupAPI(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetGroupApis() groupAPIInterface {
	return a.groupAPIInterface
}
