package repositories

import (
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	"gorm.io/gorm"
)

//PermissionRepoI 权限管理的Repo
type PermissionRepoInterface interface {
	AddPermissionRepo(permission models.Permission) error
	GetPermissionByIDRepo(id int) (permission models.Permission, err error)
	AddMenuRepo(menu models.Menu) error
	UpdateMenuByIdRepo(menu models.Menu) error
	UpdatePermissionByIDRepo(permission models.Permission) error
	DeletePermissionByIDRepo(id int) error
	GetMenuListRepo(menu models.Menu) ([]models.Menu, error)
	GetMenuByID(id int64) (models.Menu, error)
}

type permissionRepo struct {
	*gorm.DB
}

func NewPermissionRepo(db *gorm.DB) PermissionRepoInterface {
	return &permissionRepo{
		DB: db,
	}
}

func (p *permissionRepo) AddPermissionRepo(permission models.Permission) error {
	return p.Create(&permission).Error
}

func (p *permissionRepo) GetPermissionByIDRepo(id int) (permission models.Permission, err error) {
	err = p.First(&permission, id).Error
	return
}

func (p *permissionRepo) AddMenuRepo(menu models.Menu) error {
	menuRecord, err := p.GetMenuByNameAndModule(menu.Name, menu.Module)
	if err != nil && menuRecord.ID == 0 {
		return p.Create(&menu).Error
	}
	return errors.New("menu is exist")
}

func (p *permissionRepo) UpdateMenuByIdRepo(menu models.Menu) error {
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

func (p *permissionRepo) UpdatePermissionByIDRepo(permission models.Permission) error {
	return p.Model(&permission).Updates(permission).Error
}

func (p *permissionRepo) DeletePermissionByIDRepo(id int) error {
	var (
		permission models.Permission
	)
	if id != 0 {
		permission.ID = id
		return p.Delete(&permission).Error
	}
	return nil
}

func (p *permissionRepo) GetMenuListRepo(menu models.Menu) ([]models.Menu, error) {
	var (
		menus []models.Menu
	)
	db := p.Table("menu").
		Select("name, parent_id, module, template_path, id").
		Where("status =?", 1)
	if menu.Module != 0 {
		db = db.Where("module=?", menu.Module)
	}

	if menu.Name != "" {
		db = db.Where("name like ?", "%" + menu.Name + "%")
	}

	if menu.ID != 0 {
		db = db.Where("parent_id=?", menu.ID)
	}
	err := db.Order("`order` asc").Find(&menus).Error
	return menus, err
}

func (p *permissionRepo) GetMenuByID(id int64) (menu models.Menu, err error) {
	err = p.First(&menu, id).Error
	return
}



