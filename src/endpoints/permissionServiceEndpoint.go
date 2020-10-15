package endpoints

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
)

type PermissionServiceEndpoint struct {
	AddPermissionEndpoint          endpoint.Endpoint
	GetPermissionByIDEndpoint      endpoint.Endpoint
	UpdatePermissionByIDEndpoint   endpoint.Endpoint
	AddMenuEndpoint                endpoint.Endpoint
	UpdateMenuByIDEndpoint         endpoint.Endpoint
	DeletePermissionByIDEndpoint   endpoint.Endpoint
	GetMenuCascadeByModuleEndpoint endpoint.Endpoint
}

func NewPermissionEndpoint(service services.ServiceI) *PermissionServiceEndpoint {
	return &PermissionServiceEndpoint{
		AddPermissionEndpoint: MakeAddPermissionEndpoint(service.GetPermissionService()),
	}
}

func MakeAddPermissionEndpoint(permissionService services.PermissionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		permission, ok := request.(models.Permission)
		if !ok {
			return nil, RequestParamsTypeError
		}
		response, err = permissionService.AddPermissionSvc(ctx, permission)
		return
	}
}

func (p *PermissionServiceEndpoint) AddPermissionSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	res, err := p.AddPermissionEndpoint(ctx, permission)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (p *PermissionServiceEndpoint) GetPermissionByIDSvc(ctx context.Context, id int) (models.Permission, error) {
	res, err := p.GetPermissionByIDEndpoint(ctx, id)
	if err != nil {
		return models.Permission{}, err
	}
	return res.(models.Permission), nil
}

func (p *PermissionServiceEndpoint) UpdatePermissionByIDSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	res, err := p.UpdatePermissionByIDEndpoint(ctx, permission)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (p *PermissionServiceEndpoint) AddMenuSvc(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error) {
	res, err := p.AddMenuEndpoint(ctx, menu)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (p *PermissionServiceEndpoint) UpdateMenuByIDSvc(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error) {
	res, err := p.UpdateMenuByIDEndpoint(ctx, menu)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), err
}

func (p *PermissionServiceEndpoint) DeletePermissionByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	res, err := p.DeletePermissionByIDEndpoint(ctx, id)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return res.(pb_user_v1.NullResponse), nil
}

func (p *PermissionServiceEndpoint) GetMenuCascadeByModuleSvc(ctx context.Context, module models.MenuModule) ([]models.Cascade, error) {
	res, err := p.GetMenuCascadeByModuleEndpoint(ctx, module)
	if err != nil {
		return nil, err
	}
	return res.([]models.Cascade), nil
}
