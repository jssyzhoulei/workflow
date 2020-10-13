package endpoints

import (
	"context"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)
// GroupServiceEndpoint 组服务端点,提供给 transport 层调用
type GroupServiceEndpoint struct {
	// 声明 Group 下的所有接口
	GroupAddEndpoint endpoint.Endpoint
}

// NewGroupEndpoint GroupServiceEndpoint的构造函数
func NewGroupEndpoint(service services.ServiceI) *GroupServiceEndpoint {
	var groupServiceEndpoint = &GroupServiceEndpoint{}
	groupServiceEndpoint.GroupAddEndpoint = MakeGroupAddEndpoint(service.GetGroupService())

	return groupServiceEndpoint
}

// MakeGroupAddEndpoint 创建添加组端点,把服务包装成 Endpoint, 传入 group 接口
func MakeGroupAddEndpoint(services.GroupServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(pb_user_v1.GroupAddRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = services.GroupService{}.GroupAdd(ctx, &data)
		return
	}
}

// GroupAdd ...
func (g *GroupServiceEndpoint) GroupAdd(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
	resp, err := g.GroupAddEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), nil
}

