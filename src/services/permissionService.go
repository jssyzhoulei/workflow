package services

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/engine"
)

// PermissionServiceI 权限管理服务
type PermissionServiceInterface interface {
	//AddPermission: 新增权限
	AddPermissionSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse,error)
	//GetPermissionById: 根据id查找权限详情
	GetPermissionByIDSvc(ctx context.Context, id int) (models.Permission,error)
	//UpdatePermissionById: 根据id对权限修改
	UpdatePermissionByIDSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse,error)
	//AddMenu: 新增菜单
	AddMenuSvc(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error)
	//UpdateMenuByID: 通过id对菜单更改
	UpdateMenuByIDSvc(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error)
	//DeletePermissionByID: 删除权限通过ID
	DeletePermissionByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error)
	//GetMenuCascade: 获取菜单级联数据
	GetMenuCascadeByModuleSvc(ctx context.Context, module models.MenuModule) ([]models.Cascade, error)

}

type permissionService struct {
	permissionRepo repositories.PermissionRepoInterface
}

func NewPermissionService(e *engine.Engine) PermissionServiceInterface {
	return &permissionService{
		permissionRepo: repositories.NewPermissionRepo(e.DB.DB),
	}
}

//
func (p *permissionService) AddPermissionSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	var (
		err error
		menu models.Menu
	)
	if permission.MenuID != 0 {
		menu, err = p.permissionRepo.GetMenuByID(permission.MenuID)
		permission.Module = menu.Module
	} else {
		return pb_user_v1.NullResponse{}, errors.New("not find menu")
	}
	err = p.permissionRepo.AddPermissionRepo(permission)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) GetPermissionByIDSvc(ctx context.Context, id int) (models.Permission, error) {
	var (
		permission models.Permission
		err error
	)
	permission, err = p.permissionRepo.GetPermissionByIDRepo(id)
	return permission, err
}

func (p *permissionService) UpdatePermissionByIDSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	err := p.permissionRepo.UpdatePermissionByIDRepo(permission)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) AddMenuSvc(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = p.permissionRepo.AddMenuRepo(menu)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) UpdateMenuByIDSvc(ctx context.Context, menu models.Menu) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = p.permissionRepo.UpdateMenuByIdRepo(menu)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) DeletePermissionByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = p.permissionRepo.DeletePermissionByIDRepo(id)
	return pb_user_v1.NullResponse{}, err
}

func (p *permissionService) GetMenuCascadeByModuleSvc(ctx context.Context, module models.MenuModule) (cascades []models.Cascade,err error) {
	var (
		menu models.Menu
		menus []models.Menu
		permission models.Permission
		permissions []models.Permission
	)
	if module == 0 {
		err = errors.New("params is err of module")
		return
	}
	menu.Module = module
	permission.Module = module
	//TODO 获取权限
	permissions = permissions
	menus,err = p.permissionRepo.GetMenuListRepo(menu)
	if err != nil {
		return nil, err
	}

	cascades = menu.GetMenuPermissionCascade(menus, 0)
	return
}


