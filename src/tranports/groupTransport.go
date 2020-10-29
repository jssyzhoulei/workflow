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
	groupTreeQuery transport.Handler
	groupDelete transport.Handler
	queryGroupAndSubGroupsUsers transport.Handler
	setGroupQuotaUsed transport.Handler
	queryGroupIDAndSubGroupsID transport.Handler
}

// NewGroupGrpcTransport ...
func NewGroupGrpcTransport(endpoint *endpoints.GroupServiceEndpoint) *groupGrpcTransport {
	return &groupGrpcTransport{
		groupAdd: transport.NewServer(endpoint.GroupAddEndpoint, parser.DecodeGroupAddProto, parser.EncodeGroupProto),
		groupQueryWithQuotaByCondition: transport.NewServer(endpoint.GroupQueryWithQuotaByConditionEndpoint, parser.DecodeGroupQueryByConditionProto, parser.EncodeGroupQueryByConditionProto),
		groupUpdate: transport.NewServer(endpoint.GroupUpdateEndpoint, parser.DecodeGroupUpdateProto, parser.EncodeGroupProto),
		quotaUpdate: transport.NewServer(endpoint.QuotaUpdateEndpoint, parser.DecodeQuotaUpdateProto, parser.EncodeGroupProto),
		groupTreeQuery: transport.NewServer(endpoint.GroupTreeQueryEndpoint, parser.DecodeGroupIDProto, parser.EncodeGroupTreeQueryProto),
		groupDelete: transport.NewServer(endpoint.GroupDeleteEndpoint, parser.DecodeGroupIDProto, parser.EncodeGroupProto),
		queryGroupAndSubGroupsUsers: transport.NewServer(endpoint.QueryGroupAndSubGroupsUsersEndpoint, parser.DecodeGroupIDProto, parser.EncodeUsers),
		setGroupQuotaUsed: transport.NewServer(endpoint.SetGroupQuotaUsedEndpoint, parser.DecodeSetGroupQuotaUsedProto, parser.EncodeGroupProto),
		queryGroupIDAndSubGroupsID: transport.NewServer(endpoint.QueryGroupIDAndSubGroupsIDEndpoint, parser.DecodeGroupIDProto, parser.EncodeGroupIDsResponse),
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

// RPCGroupTreeQuery ...
func (g *groupGrpcTransport) RPCGroupTreeQuery(ctx context.Context, proto *pb_user_v1.GroupID) (*pb_user_v1.GroupTreeResponse, error) {
	_, resp, err := g.groupTreeQuery.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupTreeResponse), err
}

// RPCGroupDelete ...
func (g *groupGrpcTransport) RPCGroupDelete(ctx context.Context, proto *pb_user_v1.GroupID) (*pb_user_v1.GroupResponse, error) {
	_, resp, err := g.groupDelete.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), err
}

// RPCQueryGroupAndSubGroupsUsers ...
func (g *groupGrpcTransport) RPCQueryGroupAndSubGroupsUsers(ctx context.Context, proto *pb_user_v1.GroupID) (*pb_user_v1.Users, error) {
	_, resp, err := g.queryGroupAndSubGroupsUsers.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.Users), err
}

// RPCSetGroupQuotaUsed ...
func (g *groupGrpcTransport) RPCSetGroupQuotaUsed(ctx context.Context, proto *pb_user_v1.SetGroupQuotaUsedRequest) (*pb_user_v1.GroupResponse, error) {
	_, resp, err := g.setGroupQuotaUsed.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), err
}

// RPCSetGroupQuotaUsed ...
func (g *groupGrpcTransport) RPCQueryGroupIDAndSubGroupsID(ctx context.Context, proto *pb_user_v1.GroupID) (*pb_user_v1.GroupIDsResponse, error) {
	_, resp, err := g.queryGroupIDAndSubGroupsID.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupIDsResponse), err
}
