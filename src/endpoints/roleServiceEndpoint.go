package endpoints

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

type RoleServiceEndpoint struct {
	AddRoleEndpoint    endpoint.Endpoint
	UpdateRoleEndpoint endpoint.Endpoint
	DeleteRoleEndpoint endpoint.Endpoint
}

func NewRoleEndpoint(service services.ServiceI) *RoleServiceEndpoint {
	var roleServiceEndpoint = &RoleServiceEndpoint{}
	roleServiceEndpoint.AddRoleEndpoint = MakeAddRoleEndpoint(service.GetRoleService())
	roleServiceEndpoint.UpdateRoleEndpoint = MakeUpdateRoleEndpoint(service.GetRoleService())
	roleServiceEndpoint.DeleteRoleEndpoint = MakeDeleteRoleEndpoint(service.GetRoleService())
	return roleServiceEndpoint
}

func MakeAddRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		err = roleService.AddRoleSvc(ctx, role)
		return
	}
}

func MakeUpdateRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		err = roleService.UpdateRoleSvc(ctx, role)
		return
	}
}

func MakeDeleteRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		err = roleService.DeleteRoleSvc(ctx, role)
		return
	}
}

func (r *RoleServiceEndpoint) AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error {
	_, err := r.AddRoleEndpoint(ctx, role)
	return err
}

func (r *RoleServiceEndpoint) UpdateRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error {
	_, err := r.UpdateRoleEndpoint(ctx, role)
	return err
}

func (r *RoleServiceEndpoint) DeleteRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error {
	_, err := r.DeleteRoleEndpoint(ctx, role)
	return err
}