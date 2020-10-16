package tranports

import (
	"gitee.com/grandeep/org-svc/src/endpoints"
	"context"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

// OrgTransport 实现了 pb_user_v1.RpcOrgServiceServer
type OrgTransport struct {
	*userGrpcTransport
	*groupGrpcTransport
	*permissionGrpcTransport
	*roleGrpcTransport
}


func (o *OrgTransport) RpcUpdateUser(c context.Context, proto *pb_user_v1.UserProto) (*pb_user_v1.NullResponse, error) {
	panic("implement me")
}

//func (o *OrgTransport) RpcAddRole(ctx context.Context, proto *pb_user_v1.CreateMenuPermRequestProto) (*pb_user_v1.RoleProto, error) {
//	panic("implement me")
//}

func NewTransport(ept *endpoints.OrgEndpoint) pb_user_v1.RpcOrgServiceServer {
	return &OrgTransport{
		userGrpcTransport: NewUserGrpcTransport(ept.UserServiceEndpoint),
		groupGrpcTransport: NewGroupGrpcTransport(ept.GroupServiceEndpoint),
		permissionGrpcTransport: NewPermissionGrpcTransport(ept.PermissionServiceEndpoint),
		roleGrpcTransport: NewRoleGrpcTransport(ept.RoleServiceEndpoint),
	}
}
