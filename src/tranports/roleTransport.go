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
}

func NewRoleGrpcTransport(endpoint *endpoints.RoleServiceEndpoint) *roleGrpcTransport {
	var (
		addRoleServer    = transport.NewServer(endpoint.AddRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeRoleProto)
		updateRoleServer = transport.NewServer(endpoint.UpdateRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeRoleProto)
		deleteRoleServer = transport.NewServer(endpoint.DeleteRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeRoleProto)
	)
	return &roleGrpcTransport{
		addRole:    addRoleServer,
		updateRole: updateRoleServer,
		deleteRole: deleteRoleServer,
	}
}

func (u *roleGrpcTransport) RpcAddRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.UserProto, error) {
	_, user, err := u.addRole.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return user.(*pb_user_v1.UserProto), err
}

func (u *roleGrpcTransport) RpcUpdateRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.UserProto, error) {
	_, user, err := u.updateRole.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return user.(*pb_user_v1.UserProto), err
}

func (u *roleGrpcTransport) RpcDeleteRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.UserProto, error) {
	_, user, err := u.deleteRole.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return user.(*pb_user_v1.UserProto), err
}
