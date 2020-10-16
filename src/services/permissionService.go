package services

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
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
	GetMenuCascadeByModuleSvc(ctx context.Context, module models.MenuModule) (*pb_user_v1.Cascades, error)
}

type permissionService struct {
	permissionRepo repositories.PermissionRepoInterface
}

func NewPermissionService(repo repositories.RepoI) PermissionServiceInterface {
	return &permissionService{
		permissionRepo: repo.GetPermissionRepo(),
	}
}

//
func (p *permissionService) AddPermissionSvc(ctx context.Context, permission models.Permission) (pb_user_v1.NullResponse, error) {
	var (
		err error
		menu models.Menu
	)
	if permission.MenuID != 0 {
		menu, err = p.permissionRepo.GetMenuByIDRepo(permission.MenuID)
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

func (p *permissionService) GetMenuCascadeByModuleSvc(ctx context.Context, module models.MenuModule) (c *pb_user_v1.Cascades,err error) {
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
	permissions,err = p.permissionRepo.GetPermissionListRepo(permission, nil)
	if err != nil {
		return c, err
	}
	menus,err = p.permissionRepo.GetMenuListRepo(menu)
	if err != nil {
		return c, err
	}

	cascades := GetMenuCascade(menus, -1)
	cascades = AddPermissionCascade(permissions, cascades)
	c = &pb_user_v1.Cascades{
		Cascades:             cascades,
	}
	return
}


func GetMenuCascade(menus []models.Menu, parentId int) (cascades []*pb_user_v1.Cascade) {
	for k, menu := range menus {
		var (
			cascade = &pb_user_v1.Cascade{}
		)
		if parentId == menu.ParentID {
			var (
				menusNew = make([]models.Menu, len(menus)-1)
			)
			copy(menusNew[:k], menus[:k])
			copy(menusNew[k:], menus[k+1:])
			cascade.Value = int64(menu.ID)
			cascade.Label = menu.Name
			cascade.Children = GetMenuCascade(menusNew, menu.ID)
			cascades = append(cascades, cascade)
		}
	}
	return cascades
}

func AddPermissionCascade(permissions []models.Permission, cascades []*pb_user_v1.Cascade) []*pb_user_v1.Cascade {
	for k, cascade := range cascades {
		if len(cascade.Children) > 0 {
			cs := AddPermissionCascade(permissions, cascade.Children)
			cascades[k].Children = cs
			continue
		} else {
			for index, permission := range permissions {
				if cascade.Value == int64(permission.MenuID) {
					var (
						c = &pb_user_v1.Cascade{}
						permissionsNew = make([]models.Permission, len(permissions) -1)
					)
					copy(permissionsNew[:index], permissions[:index])
					copy(permissionsNew[index:], permissions[index+1:])
					c.Value = int64(permission.ID)
					c.Label = permission.UriName
					cascade.Child = append(cascade.Child, c)
				}
			}
			cascades[k] = cascade
		}
	}
	return cascades
}