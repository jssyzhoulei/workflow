package endpoints

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

type UserServiceEndpoint struct {
	AddUserEndpoint endpoint.Endpoint
}

func NewUserEndpoint(service services.ServiceI) *UserServiceEndpoint {
	var userServiceEndpoint = &UserServiceEndpoint{}
	userServiceEndpoint.AddUserEndpoint = MakeAddUserEndpoint(service.GetUserService())
	return userServiceEndpoint
}

var (
	RequestParamsTypeError = errors.New("request params type error")
)

func MakeAddUserEndpoint(userService services.UserServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.User2)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.AddUserSvc(ctx, user)
		return
	}
}

func (u *UserServiceEndpoint) AddUserSvc(ctx context.Context, user models.User2) (models.User2, error) {
	res, err := u.AddUserEndpoint(ctx, user)
	if err != nil {
		return user, err
	}
	return res.(models.User2), nil
}
