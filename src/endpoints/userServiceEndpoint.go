package endpoints

import (
	"context"
	"errors"
	"fmt"
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
	BatchDeleteUsersEndpoint endpoint.Endpoint
	AddUsersEndpoint endpoint.Endpoint
	ImportUsersByGroupIdEndpoint endpoint.Endpoint
}

// NewUserEndpoint UserServiceEndpoint 的构造函数
func NewUserEndpoint(service services.ServiceI) *UserServiceEndpoint {
	var userServiceEndpoint = &UserServiceEndpoint{}
	userServiceEndpoint.AddUserEndpoint = MakeAddUserEndpoint(service.GetUserService())
	userServiceEndpoint.GetUserByIDEndpoint = MakeGetUserByIDEndpoint(service.GetUserService())
	userServiceEndpoint.UpdateUserByIDEndpoint = MakeUpdataUserByIDEndpoint(service.GetUserService())
	userServiceEndpoint.DeleteUserByIDEndpoint = MakeDeleteUserByIDEndpoint(service.GetUserService())
	userServiceEndpoint.GetUserListEndpoint = MakeGetUserListEndpoint(service.GetUserService())
	userServiceEndpoint.BatchDeleteUsersEndpoint = MakeBatchDeleteUsersEndpoint(service.GetUserService())
	userServiceEndpoint.AddUsersEndpoint = MakeAddUsersEndpoint(service.GetUserService())
	userServiceEndpoint.ImportUsersByGroupIdEndpoint = MakeImportUsersByGroupIdEndpoint(service.GetUserService())
	return userServiceEndpoint
}

func MakeAddUsersEndpoint(service services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		addUserReq, ok := request.(*pb_user_v1.AddUsersRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = service.AddUsersSvc(ctx, addUserReq)
		return
	}
}

var (
	RequestParamsTypeError = errors.New("request params type error")
)

// MakeAddUserEndpoint 创建添加用户端点，把服务包装成 Endpoint，传入 user 接口
func MakeAddUserEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userRoleDTO, ok := request.(models.UserRolesDTO)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.AddUserSvc(ctx, userRoleDTO)
		return
	}
}

// MakeGetUserByIDEndpoint ...
func MakeGetUserByIDEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(int)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.GetUserByIDSvc(ctx, id)
		return
	}
}

// MakeUpdataUserByIDEndpoint ...
func MakeUpdataUserByIDEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.UserRolesDTO)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.UpdateUserByIDSvc(ctx, user)
		return
	}
}

func MakeImportUsersByGroupIdEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(int)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.ImportUsersByGroupIdSvc(ctx, id)
		return
	}
}

// MakeDeleteUserByIDEndpoint ...
func MakeDeleteUserByIDEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(int)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.DeleteUserByIDSvc(ctx, id)
		return
	}
}

// MakeGetUserListEndpoint ...
func MakeGetUserListEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(*pb_user_v1.UserPage)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.GetUserListSvc(ctx, user)
		fmt.Println(response, err)
		return
	}
}

// MakeBatchDeleteUsersEndpoint ...
func MakeBatchDeleteUsersEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ids, ok := request.([]int64)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.BatchDeleteUsersSvc(ctx, ids)
		fmt.Println(response, err)
		return
	}
}

// AddUserSvc ...
func (u *UserServiceEndpoint) AddUserSvc(ctx context.Context, userProto models.UserRolesDTO) (pb_user_v1.NullResponse, error) {
	resp, err := u.AddUserEndpoint(ctx, userProto)
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
func (u *UserServiceEndpoint) UpdateUserByIDSvc(ctx context.Context, userProto models.UserRolesDTO) (pb_user_v1.NullResponse, error){
	resp, err := u.UpdateUserByIDEndpoint(ctx, userProto)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return resp.(pb_user_v1.NullResponse), nil
}

func (u *UserServiceEndpoint) ImportUsersByGroupIdSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error){
	resp, err := u.ImportUsersByGroupIdEndpoint(ctx, id)
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
func (u *UserServiceEndpoint) GetUserListSvc(ctx context.Context, user *pb_user_v1.UserPage) (c *pb_user_v1.UsersPage, err error) {
	resp, err := u.GetUserListEndpoint(ctx, user)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.UsersPage), nil
}

// BatchDeleteUsersSvc ...
func (u *UserServiceEndpoint) BatchDeleteUsersSvc(ctx context.Context, ids []int64) (pb_user_v1.NullResponse, error) {
	resp, err := u.BatchDeleteUsersEndpoint(ctx, ids)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return resp.(pb_user_v1.NullResponse), nil
}
func (u *UserServiceEndpoint) AddUsersSvc(ctx context.Context, users *pb_user_v1.AddUsersRequest) (pb_user_v1.NullResponse, error) {
	resp, err := u.AddUsersEndpoint(ctx, users)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return resp.(pb_user_v1.NullResponse), nil
}
