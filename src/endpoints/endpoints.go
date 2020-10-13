package endpoints

import (
	"gitee.com/grandeep/org-svc/src/services"
)

type OrgEndpoint struct {
	*UserServiceEndpoint
	*GroupServiceEndpoint
}


func NewEndpoint(service services.ServiceI) *OrgEndpoint {
	return &OrgEndpoint{
		UserServiceEndpoint: NewUserEndpoint(service),
		GroupServiceEndpoint: NewGroupEndpoint(service),
	}
}
