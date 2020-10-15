package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/tranports/parser"
	transport "github.com/go-kit/kit/transport/grpc"
)

type permissionGrpcTransport struct {
	addPermission transport.Handler
}

func NewPermissionGrpcTransport(permissionEndpoint *endpoints.PermissionServiceEndpoint) *permissionGrpcTransport {
	return &permissionGrpcTransport{
		addPermission: transport.NewServer(permissionEndpoint.AddPermissionEndpoint, parser.DecodePermissionProto, parser.EncodeNullProto),
	}
}



func (p *permissionGrpcTransport) RpcAddPermission(ctx context.Context, proto *pb_user_v1.PermissionProto) (*pb_user_v1.NullResponse, error) {
	_,res,err :=  p.addPermission.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.NullResponse), nil
}

func (p *permissionGrpcTransport) RpcGetPermissionByID(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.PermissionProto, error) {
	panic("implement me")
}

func (p *permissionGrpcTransport) RpcUpdatePermissionByID(ctx context.Context, proto *pb_user_v1.PermissionProto) (*pb_user_v1.NullResponse, error) {
	panic("implement me")
}

func (p *permissionGrpcTransport) RpcAddMenu(ctx context.Context, proto *pb_user_v1.MenuProto) (*pb_user_v1.NullResponse, error) {
	panic("implement me")
}

func (p *permissionGrpcTransport) RpcUpdateMenuByID(ctx context.Context, proto *pb_user_v1.MenuProto) (*pb_user_v1.NullResponse, error) {
	panic("implement me")
}

func (p *permissionGrpcTransport) RpcDeletePermissionByID(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.NullResponse, error) {
	panic("implement me")
}

func (p *permissionGrpcTransport) RpcGetMenuCascadeByModule(ctx context.Context, module *pb_user_v1.MenuModule) (*pb_user_v1.Cascades, error) {
	panic("implement me")
}



