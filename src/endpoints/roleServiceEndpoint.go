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
	QueryRoleEndpoint  endpoint.Endpoint
	QueryRolesEndpoint endpoint.Endpoint
	MenuTreeEndpoint   endpoint.Endpoint
}

func NewRoleEndpoint(service services.ServiceI) *RoleServiceEndpoint {
	var roleServiceEndpoint = &RoleServiceEndpoint{}
	roleServiceEndpoint.AddRoleEndpoint = MakeAddRoleEndpoint(service.GetRoleService())
	roleServiceEndpoint.UpdateRoleEndpoint = MakeUpdateRoleEndpoint(service.GetRoleService())
	roleServiceEndpoint.DeleteRoleEndpoint = MakeDeleteRoleEndpoint(service.GetRoleService())
	roleServiceEndpoint.QueryRoleEndpoint = MakeQueryRoleEndpoint(service.GetRoleService())
	roleServiceEndpoint.QueryRolesEndpoint = MakeQueryRolesEndpoint(service.GetRoleService())
	roleServiceEndpoint.MenuTreeEndpoint = MakeMenuTreeEndpoint(service.GetRoleService())
	return roleServiceEndpoint
}

func MakeAddRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(*models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.AddRoleSvc(ctx, role)
		return
	}
}

func MakeUpdateRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		role, ok := request.(*models.CreateMenuPermRequest)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.UpdateRoleSvc(ctx, role)
		return
	}
}

func MakeDeleteRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(int)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.DeleteRoleSvc(ctx, id)
		return
	}
}

func MakeQueryRoleEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(int)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.QueryRoleSvc(ctx, id)
		return
	}
}

func MakeQueryRolesEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		page, ok := request.(*pb_user_v1.RolePageRequestProto)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.QueryRolesSvc(ctx, page)
		return
	}
}

func MakeMenuTreeEndpoint(roleService services.RoleServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		m, ok := request.(models.MenuModule)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = roleService.MenuTreeSvc(ctx, m)
		return
	}
}

func (r *RoleServiceEndpoint) AddRoleSvc(ctx context.Context, role *models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	res, err := r.AddRoleEndpoint(ctx, role)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil

}

func (r *RoleServiceEndpoint) UpdateRoleSvc(ctx context.Context, role *models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	res, err := r.UpdateRoleEndpoint(ctx, role)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (r *RoleServiceEndpoint) DeleteRoleSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	res, err := r.DeleteRoleEndpoint(ctx, id)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (r *RoleServiceEndpoint) QueryRoleSvc(ctx context.Context, id int) (*models.CreateMenuPermRequest, error) {
	res, err := r.QueryRoleEndpoint(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.(*models.CreateMenuPermRequest), nil
}

func (r *RoleServiceEndpoint) QueryRolesSvc(ctx context.Context, page *pb_user_v1.RolePageRequestProto) (*pb_user_v1.RolePageRequestProto, error) {
	res, err := r.QueryRolesEndpoint(ctx, page)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.RolePageRequestProto), nil
}

func (r *RoleServiceEndpoint) MenuTreeSvc(ctx context.Context, m models.MenuModule) (*pb_user_v1.Cascades, error) {
	res, err := r.MenuTreeEndpoint(ctx, m)
	if err != nil {
		return nil, err
	}
	return res.(*pb_user_v1.Cascades), nil
}
