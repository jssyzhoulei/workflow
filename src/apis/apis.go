package apis

import "gitee.com/grandeep/org-svc/src/services"

type IApis interface {
	GetUserApis() userApiI
}

type apis struct {
	userApiI
}


func NewApis(userService services.UserServiceI) IApis {
	return &apis{
		userApiI: NewUserApi(userService),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}
