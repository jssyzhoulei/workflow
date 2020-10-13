package endpoints

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

type RoleServiceEndpoint struct {
	AddRoleEndpoint endpoint.Endpoint
}

func NewRoleEndpoint(service services.ServiceI) *RoleServiceEndpoint {
	var roleServiceEndpoint = &RoleServiceEndpoint{}
	roleServiceEndpoint.AddRoleEndpoint = MakeAddRoleEndpoint(service.GetRoleService())
	return roleServiceEndpoint
}

func MakeAddRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.AddRoleSvc(ctx, role)
		return
	}
}

func (r *RoleServiceEndpoint) AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) (models.Role, error) {
	res, err := r.AddRoleEndpoint(ctx, role)
	if err != nil {
		return role.Role, err
	}
	return res.(models.Role), nil
}
