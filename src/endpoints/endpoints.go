package endpoints

import (
	"gitee.com/grandeep/org-svc/src/services"
)

type OrgEndpoint struct {
	*UserServiceEndpoint
}


func NewEndpoint(service services.ServiceI) *OrgEndpoint {
	return &OrgEndpoint{
		UserServiceEndpoint: NewUserEndpoint(service),
	}
}
