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
	queryRoles transport.Handler
	menuTree   transport.Handler
}

func NewRoleGrpcTransport(endpoint *endpoints.RoleServiceEndpoint) *roleGrpcTransport {
	var (
		addRoleServer    = transport.NewServer(endpoint.AddRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeNullProto)
		updateRoleServer = transport.NewServer(endpoint.UpdateRoleEndpoint, parser.DecodeCreateMenuPermRequestProto, parser.EncodeNullProto)
		deleteRoleServer = transport.NewServer(endpoint.DeleteRoleEndpoint, parser.DecodeIndexProto, parser.EncodeNullProto)
		queryRoleServer  = transport.NewServer(endpoint.QueryRoleEndpoint, parser.DecodeIndexProto, parser.EncodeCreateMenuPermRequest)
		queryRolesServer = transport.NewServer(endpoint.QueryRolesEndpoint, parser.DecodeRolePageProto, parser.EncodeRolePageProto)
		menuTreeServer   = transport.NewServer(endpoint.MenuTreeEndpoint, parser.DecodeModuleProto, parser.EncodeCascadeProto)
	)
	return &roleGrpcTransport{
		addRole:    addRoleServer,
		updateRole: updateRoleServer,
		deleteRole: deleteRoleServer,
		queryRole:  queryRoleServer,
		queryRoles: queryRolesServer,
		menuTree:   menuTreeServer,
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

func (u *roleGrpcTransport) RpcQueryRoles(ctx context.Context, page *pb_user_v1.RolePageRequestProto) (*pb_user_v1.RolePageRequestProto, error) {
	_, role, err := u.queryRoles.ServeGRPC(ctx, page)
	if err != nil {
		return nil, err
	}
	return role.(*pb_user_v1.RolePageRequestProto), err
}

func (u *roleGrpcTransport) RpcMenuTree(ctx context.Context, module *pb_user_v1.MenuModule) (*pb_user_v1.Cascades, error) {
	_, role, err := u.menuTree.ServeGRPC(ctx, module)
	if err != nil {
		return nil, err
	}
	return role.(*pb_user_v1.Cascades), err
}
