package repositories

import (
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
)

type RoleRepoI interface {
	AddRoleRepo(role *models.CreateMenuPermRequest) error
	BatchCreateMenuPermRepo(mps *[]models.RoleMenuPermission) error
}


type roleRepo struct {
	repo *yorm.DBRepo
	*gorm.DB
}

func NewRoleRepo(db *yorm.DB) RoleRepoI {
	return &roleRepo{
		DB: db.DB,
	}
}

func (u *roleRepo) AddRoleRepo(role *models.CreateMenuPermRequest) error {
	return u.DB.Model(models.Role{}).Create(&role.Role).Error
}

func (u *roleRepo) BatchCreateMenuPermRepo(mps *[]models.RoleMenuPermission) error {
	return u.DB.Model(models.Role{}).Create(mps).Error
}
