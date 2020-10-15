package endpoints

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

// UserServiceEndpoint 用户服务端点，提供给 transport 层调用
type UserServiceEndpoint struct {
	// 声明 User 下的所有接口
	AddUserEndpoint endpoint.Endpoint
	GetUserByIDEndpoint endpoint.Endpoint
	UpdateUserByIDEndpoint endpoint.Endpoint
	DeleteUserByIDEndpoint endpoint.Endpoint
	GetUserListEndpoint endpoint.Endpoint
}

// NewUserEndpoint UserServiceEndpoint 的构造函数
func NewUserEndpoint(service services.ServiceI) *UserServiceEndpoint {
	var userServiceEndpoint = &UserServiceEndpoint{}
	userServiceEndpoint.AddUserEndpoint = MakeAddUserEndpoint(service.GetUserService())
	userServiceEndpoint.GetUserByIDEndpoint = MakeGetUserByIDEndpoint(service.GetUserService())
	userServiceEndpoint.UpdateUserByIDEndpoint = MakeUpdataUserByIDEndpoint(service.GetUserService())
	userServiceEndpoint.DeleteUserByIDEndpoint = MakeDeleteUserByIDEndpoint(service.GetUserService())
	userServiceEndpoint.GetUserListEndpoint = MakeGetUserListEndpoint(service.GetUserService())
	return userServiceEndpoint
}

var (
	RequestParamsTypeError = errors.New("request params type error")
)

// MakeAddUserEndpoint 创建添加用户端点，把服务包装成 Endpoint，传入 user 接口
func MakeAddUserEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.User)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.AddUserSvc(ctx, user)
		return
	}
}

// MakeGetUserByIDEndpoint ...
func MakeGetUserByIDEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.User)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.GetUserByIDSvc(ctx, user.ID)
		return
	}
}

// MakeUpdataUserByIDEndpoint ...
func MakeUpdataUserByIDEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.User)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.UpdateUserByIDSvc(ctx, user)
		return
	}
}

// MakeDeleteUserByIDEndpoint ...
func MakeDeleteUserByIDEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.User)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.DeleteUserByIDSvc(ctx, user.ID)
		return
	}
}

// MakeGetUserListEndpoint ...
func MakeGetUserListEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.User)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.GetUserListSvc(ctx, user, &models.Page{})
		return
	}
}
// AddUserSvc ...
func (u *UserServiceEndpoint) AddUserSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error) {
	resp, err := u.AddUserEndpoint(ctx, user)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return resp.(pb_user_v1.NullResponse), nil
}

// GetUserByIDSvc ...
func (u *UserServiceEndpoint) GetUserByIDSvc(ctx context.Context, id int) (models.User, error) {
	resp, err := u.GetUserByIDEndpoint(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return resp.(models.User), nil
}

// UpdateUserByIDSvc ...
func (u *UserServiceEndpoint) UpdateUserByIDSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error){
	resp, err := u.UpdateUserByIDEndpoint(ctx, user)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return resp.(pb_user_v1.NullResponse), nil
}

// DeleteUserByIDSvc ...
func (u *UserServiceEndpoint) DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error){
	resp, err := u.DeleteUserByIDEndpoint(ctx, id)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return resp.(pb_user_v1.NullResponse), nil
}

// GetUserListSvc ...
func (u *UserServiceEndpoint) GetUserListSvc(ctx context.Context, user models.User, page *models.Page) ([]models.User, error){
	resp, err := u.GetUserListEndpoint(ctx, user)
	if err != nil {
		return nil, err
	}
	return resp.([]models.User), nil
}