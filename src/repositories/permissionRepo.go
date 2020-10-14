package repositories

import (
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	"gorm.io/gorm"
)

//PermissionRepoI 权限管理的Repo
type PermissionRepoI interface {
	AddPermission(permission models.Permission) error
	GetPermissionByID(id int) (permission models.Permission, err error)
	AddMenu(menu models.Menu) error
	UpdateMenuById(menu models.Menu) error
	UpdatePermissionByID(permission models.Permission) error
	DeletePermissionByID(id int) error
}

type permissionRepo struct {
	*gorm.DB
}

func NewPermissionRepo(db *gorm.DB) PermissionRepoI {
	return &permissionRepo{
		DB: db,
	}
}

func (p *permissionRepo) AddPermission(permission models.Permission) error {
	return p.Create(&permission).Error
}

func (p *permissionRepo) GetPermissionByID(id int) (permission models.Permission, err error) {
	err = p.First(&permission, id).Error
	return
}

func (p *permissionRepo) AddMenu(menu models.Menu) error {
	menuRecord, err := p.GetMenuByNameAndModule(menu.Name, menu.Module)
	if err != nil && menuRecord.ID == 0 {
		return p.Create(&menu).Error
	}
	return errors.New("menu is exist")
}

func (p *permissionRepo) UpdateMenuById(menu models.Menu) error {
	menuRecord, err := p.GetMenuByNameAndModule(menu.Name, menu.Module)
	if err != nil || menuRecord.ID == menu.ID {
		return p.Model(&menu).Updates(menu).Error
	}
	return errors.New("menu is exist")
}

func (p *permissionRepo) GetMenuByNameAndModule(name string, module models.MenuModule) (models.Menu, error) {
	var (
		menu models.Menu
		err error
	)
	err = p.Where("name = ?", name).Where("module=?", module).First(&menu).Error
	return menu, err
}

func (p *permissionRepo) UpdatePermissionByID(permission models.Permission) error {
	return p.Model(&permission).Updates(permission).Error
}

func (p *permissionRepo) DeletePermissionByID(id int) error {
	var (
		permission models.Permission
	)
	if id != 0 {
		permission.ID = id
		return p.Delete(&permission).Error
	}
	return nil
}


