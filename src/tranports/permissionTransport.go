package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/tranports/parser"
	transport "github.com/go-kit/kit/transport/grpc"
)

type permissionGrpcTransport struct {
	addPermission          transport.Handler
	addMenu                transport.Handler
	getMenuCascadeByModule transport.Handler
	getPermissionByID      transport.Handler
	deletePermissionByID   transport.Handler
	updatePermissionByID   transport.Handler
}

func NewPermissionGrpcTransport(permissionEndpoint *endpoints.PermissionServiceEndpoint) *permissionGrpcTransport {
	return &permissionGrpcTransport{
		addPermission:          transport.NewServer(permissionEndpoint.AddPermissionEndpoint, parser.DecodePermissionProto, parser.EncodeNullProto),
		addMenu:                transport.NewServer(permissionEndpoint.AddMenuEndpoint, parser.DecodeMenuProto, parser.EncodeNullProto),
		getMenuCascadeByModule: transport.NewServer(permissionEndpoint.GetMenuCascadeByModuleEndpoint, parser.DecodeModuleProto, parser.EncodeCascadeProto),
		getPermissionByID:      transport.NewServer(permissionEndpoint.GetPermissionByIDEndpoint, parser.DecodeIndexProto, parser.EncodePermissionProto),
		deletePermissionByID: transport.NewServer(permissionEndpoint.DeletePermissionByIDEndpoint, parser.DecodeIndexProto, parser.EncodeNullProto),
		updatePermissionByID: transport.NewServer(permissionEndpoint.UpdatePermissionByIDEndpoint, parser.DecodePermissionProto, parser.EncodeNullProto),
	}
}

func (p *permissionGrpcTransport) RpcAddPermission(ctx context.Context, proto *pb_user_v1.PermissionProto) (*pb_user_v1.NullResponse, error) {
	_, res, err := p.addPermission.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.NullResponse), nil
}

func (p *permissionGrpcTransport) RpcGetPermissionByID(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.PermissionProto, error) {
	_, res, err := p.getPermissionByID.ServeGRPC(ctx, index)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.PermissionProto), nil
}

func (p *permissionGrpcTransport) RpcUpdatePermissionByID(ctx context.Context, proto *pb_user_v1.PermissionProto) (*pb_user_v1.NullResponse, error) {
	_, res, err := p.updatePermissionByID.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.NullResponse), nil
}

func (p *permissionGrpcTransport) RpcAddMenu(ctx context.Context, proto *pb_user_v1.MenuProto) (*pb_user_v1.NullResponse, error) {
	_, res, err := p.addMenu.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.NullResponse), nil
}

func (p *permissionGrpcTransport) RpcUpdateMenuByID(ctx context.Context, proto *pb_user_v1.MenuProto) (*pb_user_v1.NullResponse, error) {
	panic("implement me")
}

func (p *permissionGrpcTransport) RpcDeletePermissionByID(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.NullResponse, error) {
	_, res, err := p.deletePermissionByID.ServeGRPC(ctx, index)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.NullResponse), nil
}

func (p *permissionGrpcTransport) RpcGetMenuCascadeByModule(ctx context.Context, module *pb_user_v1.MenuModule) (*pb_user_v1.Cascades, error) {
	_, res, err := p.getMenuCascadeByModule.ServeGRPC(ctx, module)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.Cascades), nil
}
