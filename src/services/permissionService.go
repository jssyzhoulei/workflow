package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/engine"
)

// PermissionServiceI 权限管理服务
type PermissionServiceI interface {
	//AddPermission: 新增权限
	AddPermission(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse,error)
	//GetPermissionById: 根据id查找权限详情
	GetPermissionByID(ctx context.Context, id int) (models.Permission,error)
	//UpdatePermissionById: 根据id对权限修改
	UpdatePermissionByID(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse,error)
	//AddMenu: 新增菜单
	AddMenu(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error)
	//UpdateMenuByID: 通过id对菜单更改
	UpdateMenuByID(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error)

	DeletePermissionByID(ctx context.Context, id int) (pb_user_v1.NullResponse, error)
}

type permissionService struct {
	permissionRepo repositories.PermissionRepoI
}

func NewPermissionService(e *engine.Engine) PermissionServiceI {
	return &permissionService{
		permissionRepo: repositories.NewPermissionRepo(e.DB.DB),
	}
}

//
func (p *permissionService) AddPermission(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	err := p.permissionRepo.AddPermission(permission)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) GetPermissionByID(ctx context.Context, id int) (models.Permission, error) {
	var (
		permission models.Permission
		err error
	)
	permission, err = p.permissionRepo.GetPermissionByID(id)
	return permission, err
}

func (p *permissionService) UpdatePermissionByID(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	err := p.permissionRepo.UpdatePermissionByID(permission)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) AddMenu(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = p.permissionRepo.AddMenu(menu)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) UpdateMenuByID(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = p.permissionRepo.UpdateMenuById(menu)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) DeletePermissionByID(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = p.permissionRepo.DeletePermissionByID(id)
	return pb_user_v1.NullResponse{}, err
}

