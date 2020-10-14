package endpoints

import (
	"context"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

// GroupServiceEndpoint 组服务端点结构体实现 services.GroupServiceInterface
// 需要添加所有接口的 Endpoint 到此结构体
type GroupServiceEndpoint struct {
	// GroupAddEndpoint 添加组
	GroupAddEndpoint endpoint.Endpoint
}

// NewGroupEndpoint GroupServiceEndpoint的构造函数
func NewGroupEndpoint(service services.ServiceI) *GroupServiceEndpoint {
	return &GroupServiceEndpoint{
		GroupAddEndpoint: MakeGroupAddEndpoint(service.GetGroupService()),
	}
}

// MakeGroupAddEndpoint 把服务包装成 Endpoint, 在 NewGroupEndpoint 中使用
func MakeGroupAddEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.GroupAddRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.GroupAddSvc(ctx, data)
		return
	}
}

// GroupAdd ...
func (g *GroupServiceEndpoint) GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
	resp, err := g.GroupAddEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), nil
}
