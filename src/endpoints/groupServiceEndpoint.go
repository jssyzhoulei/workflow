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
	GroupQueryWithQuotaByConditionEndpoint endpoint.Endpoint
	GroupUpdateEndpoint endpoint.Endpoint
	QuotaUpdateEndpoint endpoint.Endpoint
	GroupTreeQueryEndpoint endpoint.Endpoint
	GroupDeleteEndpoint endpoint.Endpoint
	QueryGroupAndSubGroupsUsersEndpoint endpoint.Endpoint
}

// NewGroupEndpoint GroupServiceEndpoint的构造函数
func NewGroupEndpoint(service services.ServiceI) *GroupServiceEndpoint {
	return &GroupServiceEndpoint{
		GroupAddEndpoint: MakeGroupAddEndpoint(service.GetGroupService()),
		GroupQueryWithQuotaByConditionEndpoint: MakeGroupQueryWithQuotaByConditionEndpoint(service.GetGroupService()),
		GroupUpdateEndpoint: MakeGroupUpdateEndpoint(service.GetGroupService()),
		QuotaUpdateEndpoint: MakeQuotaUpdateEndpoint(service.GetGroupService()),
		GroupTreeQueryEndpoint: MakeGroupTreeQueryEndpoint(service.GetGroupService()),
		GroupDeleteEndpoint: MakeGroupDeleteEndpoint(service.GetGroupService()),
		QueryGroupAndSubGroupsUsersEndpoint: MakeQueryGroupAndSubGroupsUsersEndpoint(service.GetGroupService()),
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

// GroupAddSvc ...
func (g *GroupServiceEndpoint) GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
	resp, err := g.GroupAddEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), nil
}

// MakeGroupQueryWithQuotaByConditionEndpoint ...
func MakeGroupQueryWithQuotaByConditionEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.GroupQueryWithQuotaByConditionRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.GroupQueryWithQuotaByConditionSvc(ctx, data)
		return
	}
}

// GroupQueryWithQuotaByConditionSvc ...
func (g *GroupServiceEndpoint) GroupQueryWithQuotaByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error) {

	resp , err := g.GroupQueryWithQuotaByConditionEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupQueryWithQuotaByConditionResponse), nil
}

// MakeGroupUpdateEndpoint ...
func MakeGroupUpdateEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.GroupUpdateRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.GroupUpdateSvc(ctx, data)
		return
	}
}

// GroupUpdateSvc ...
func (g *GroupServiceEndpoint) GroupUpdateSvc(ctx context.Context, data *pb_user_v1.GroupUpdateRequest) (*pb_user_v1.GroupResponse, error) {

	resp , err := g.GroupUpdateEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), nil
}

// MakeQuotaUpdateEndpoint ...
func MakeQuotaUpdateEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.QuotaUpdateRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.QuotaUpdateSvc(ctx, data)
		return
	}
}

// QuotaUpdateSvc ...
func (g *GroupServiceEndpoint) QuotaUpdateSvc(ctx context.Context, data *pb_user_v1.QuotaUpdateRequest) (*pb_user_v1.GroupResponse, error) {

	resp , err := g.QuotaUpdateEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), nil
}

// MakeGroupTreeQueryEndpoint ...
func MakeGroupTreeQueryEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.GroupID)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.GroupTreeQuerySvc(ctx, data)
		return
	}
}

// GroupTreeQuerySvc ...
func (g *GroupServiceEndpoint) GroupTreeQuerySvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupTreeResponse, error) {

	resp , err := g.GroupTreeQueryEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupTreeResponse), nil
}

// MakeGroupDeleteEndpoint ...
func MakeGroupDeleteEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.GroupID)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.GroupDeleteSvc(ctx, data)
		return
	}
}

// GroupDeleteSvc ...
func (g *GroupServiceEndpoint) GroupDeleteSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupResponse, error) {

	resp , err := g.GroupDeleteEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), nil
}

// MakeQueryGroupAndSubGroupsUsersEndpoint ...
func MakeQueryGroupAndSubGroupsUsersEndpoint(groupServiceInterface services.GroupServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(*pb_user_v1.GroupID)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = groupServiceInterface.QueryGroupAndSubGroupsUsersSvc(ctx, data)
		return
	}
}

// QueryGroupAndSubGroupsUsersSvc ...
func (g *GroupServiceEndpoint) QueryGroupAndSubGroupsUsersSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error) {
	resp , err := g.QueryGroupAndSubGroupsUsersEndpoint(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.Users), nil
}