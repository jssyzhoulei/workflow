package endpoints

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
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
		response, err = roleService.AddRoleSvc(ctx, role)
		return
	}
}

func MakeUpdateRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.UpdateRoleSvc(ctx, role)
		return
	}
}

func MakeDeleteRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.DeleteRoleSvc(ctx, role)
		return
	}
}

func (r *RoleServiceEndpoint) AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	res, err :=  r.AddRoleEndpoint(ctx, role)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil

}

func (r *RoleServiceEndpoint) UpdateRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	res, err :=  r.UpdateRoleEndpoint(ctx, role)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (r *RoleServiceEndpoint) DeleteRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	res, err :=  r.DeleteRoleEndpoint(ctx, role)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}