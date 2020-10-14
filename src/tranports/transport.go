package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

// OrgTransport 实现了 pb_user_v1.RpcOrgServiceServer
type OrgTransport struct {
	*userGrpcTransport
	*groupGrpcTransport
}

func (o *OrgTransport) RpcAddRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.RoleProto, error) {
	panic("implement me")
}

func NewTransport(ept *endpoints.OrgEndpoint) pb_user_v1.RpcOrgServiceServer {
	return &OrgTransport{
		userGrpcTransport: NewUserGrpcTransport(ept.UserServiceEndpoint),
		groupGrpcTransport: NewGroupGrpcTransport(ept.GroupServiceEndpoint),
	}
}
