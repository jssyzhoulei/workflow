package tranports

import "gitee.com/grandeep/org-svc/src/endpoints"

type OrgTransport struct {
	*userGrpcTransport
}

func NewTransport(ept *endpoints.OrgEndpoint) *OrgTransport {
	return &OrgTransport{
		userGrpcTransport: NewUserGrpcTransport(ept.UserServiceEndpoint),
	}
}
