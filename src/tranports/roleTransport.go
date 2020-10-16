package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/tranports/parser"
	transport "github.com/go-kit/kit/transport/grpc"
)

type roleGrpcTransport struct {
	addRole    transport.Handler
	updateRole transport.Handler
	deleteRole transport.Handler
	queryRole  transport.Handler
}

func NewRoleGrpcTransport(endpoint *endpoints.RoleServiceEndpoint) *roleGrpcTransport {
	var (
		addRoleServer    = transport.NewServer(endpoint.AddRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeNullProto)
		updateRoleServer = transport.NewServer(endpoint.UpdateRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeNullProto)
		deleteRoleServer = transport.NewServer(endpoint.DeleteRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeNullProto)
	)
	return &roleGrpcTransport{
		addRole:    addRoleServer,
		updateRole: updateRoleServer,
		deleteRole: deleteRoleServer,
	}
}

func (u *roleGrpcTransport) RpcAddRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.NullResponse, error) {
	_, role, err := u.addRole.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return role.(*pb_user_v1.NullResponse), err
}

func (u *roleGrpcTransport) RpcUpdateRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.NullResponse, error) {
	_, role, err := u.updateRole.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return role.(*pb_user_v1.NullResponse), err
}

func (u *roleGrpcTransport) RpcDeleteRole(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.NullResponse, error) {
	_, role, err := u.deleteRole.ServeGRPC(ctx, index)
	if err != nil {
		return nil, err
	}
	return role.(*pb_user_v1.NullResponse), err
}

func (u *roleGrpcTransport) RpcQueryRole(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.CreateMenuPermRequestProto, error) {
	_, role, err := u.queryRole.ServeGRPC(ctx, index)
	if err != nil {
		return nil, err
	}
	return role.(*pb_user_v1.CreateMenuPermRequestProto), err
}
