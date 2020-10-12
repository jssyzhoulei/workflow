package apis

import (
	"gitee.com/grandeep/org-svc/client"
)

type IApis interface {
	GetUserApis() userApiI
}

type apis struct {
	userApiI
}


func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}
