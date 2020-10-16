package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/tranports/parser"
	transport "github.com/go-kit/kit/transport/grpc"
)

type groupGrpcTransport struct {
	groupAdd   transport.Handler
	groupQueryWithQuotaByCondition transport.Handler
	groupUpdate transport.Handler
	quotaUpdate transport.Handler
}

// NewGroupGrpcTransport ...
func NewGroupGrpcTransport(endpoint *endpoints.GroupServiceEndpoint) *groupGrpcTransport {
	return &groupGrpcTransport{
		groupAdd: transport.NewServer(endpoint.GroupAddEndpoint, parser.DecodeGroupAddProto, parser.EncodeGroupProto),
		groupQueryWithQuotaByCondition: transport.NewServer(endpoint.GroupQueryWithQuotaByConditionEndpoint, parser.DecodeGroupQueryByConditionProto, parser.EncodeGroupQueryByConditionProto),
		groupUpdate: transport.NewServer(endpoint.GroupUpdateEndpoint, parser.DecodeGroupUpdateProto, parser.EncodeGroupProto),
		quotaUpdate: transport.NewServer(endpoint.QuotaUpdateEndpoint, parser.DecodeQuotaUpdateProto, parser.EncodeGroupProto),
	}
}

// RPCGroupAdd ...
func (g *groupGrpcTransport) RPCGroupAdd(ctx context.Context, proto *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
	_, resp, err := g.groupAdd.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), err
}

// RPCGroupQueryWithQuotaByCondition ...
func (g *groupGrpcTransport) RPCGroupQueryWithQuotaByCondition(ctx context.Context, proto *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error) {
	_, resp, err := g.groupQueryWithQuotaByCondition.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupQueryWithQuotaByConditionResponse), err
}

// RPCGroupUpdate ...
func (g *groupGrpcTransport) RPCGroupUpdate(ctx context.Context, proto *pb_user_v1.GroupUpdateRequest) (*pb_user_v1.GroupResponse, error) {
	_, resp, err := g.groupUpdate.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), err
}

// RPCQuotaUpdate ...
func (g *groupGrpcTransport) RPCQuotaUpdate(ctx context.Context, proto *pb_user_v1.QuotaUpdateRequest) (*pb_user_v1.GroupResponse, error) {
	_, resp, err := g.quotaUpdate.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), err
}