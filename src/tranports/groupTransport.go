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
}

func NewGroupGrpcTransport(endpoint *endpoints.GroupServiceEndpoint) *groupGrpcTransport {
	var (
		groupAddServer = transport.NewServer(endpoint.GroupAddEndpoint, parser.DecodeGroupAddProto, parser.EncodeGroupProto)
	)
	return &groupGrpcTransport{
		groupAdd: groupAddServer,
	}
}


func (g *groupGrpcTransport) RPCGroupAdd(ctx context.Context, proto *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
	_, resp, err := g.groupAdd.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.GroupResponse), err
}

