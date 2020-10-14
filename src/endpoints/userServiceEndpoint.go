package endpoints

import (
	"context"
	"errors"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

// UserServiceEndpoint 用户服务端点，提供给 transport 层调用
type UserServiceEndpoint struct {
	// 声明 User 下的所有接口
	AddUserEndpoint endpoint.Endpoint
}

// NewUserEndpoint UserServiceEndpoint 的构造函数
func NewUserEndpoint(service services.ServiceI) *UserServiceEndpoint {
	var userServiceEndpoint = &UserServiceEndpoint{}
	userServiceEndpoint.AddUserEndpoint = MakeAddUserEndpoint(service.GetUserService())
	return userServiceEndpoint
}

var (
	RequestParamsTypeError = errors.New("request params type error")
)

// MakeAddUserEndpoint 创建添加用户端点，把服务包装成 Endpoint，传入 group 接口
func MakeAddUserEndpoint(userService services.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(pb_user_v1.UserProto)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = userService.AddUserSvc(ctx, &user)
		return
	}
}

// AddUserSvc ...
func (u *UserServiceEndpoint) AddUserSvc(ctx context.Context, user *pb_user_v1.UserProto) (*pb_user_v1.NullResponse, error) {
	resp, err := u.AddUserEndpoint(ctx, user)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}
