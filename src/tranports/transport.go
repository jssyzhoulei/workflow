package tranports

import "gitee.com/grandeep/org-svc/src/endpoints"

type OrgTransport struct {
	*userGrpcTransport
	*groupGrpcTransport
	*roleGrpcTransport
}

func NewTransport(ept *endpoints.OrgEndpoint) *OrgTransport {
	return &OrgTransport{
		userGrpcTransport: NewUserGrpcTransport(ept.UserServiceEndpoint),
		groupGrpcTransport: NewGroupGrpcTransport(ept.GroupServiceEndpoint),
		roleGrpcTransport: NewRoleGrpcTransport(ept.RoleServiceEndpoint),
	}
}
