package repositories

import (
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
)

type RoleRepoI interface {
	GetTx() *gorm.DB
	AddRoleRepo(role *models.CreateMenuPermRequest) error
	BatchCreateMenuPermRepo(mps *[]*models.RoleMenuPermission) error
	UpdateRoleRepo(role *models.Role) error
	DeleteRoleRepo(role *models.Role) error
	ListRoleRepo(page, perPage, userId int) (*[]models.Role, error)
	RoleDetailRepo(roleId, userId int) (*models.CreateMenuPermRequest, error)
	DeleteMenuPermissionByRoleIDRepo(roleId int) error
}

type roleRepo struct {
	*gorm.DB
}

func NewRoleRepo(db *yorm.DB) RoleRepoI {
	return &roleRepo{
		DB: db.DB,
	}
}

func (u *roleRepo) GetTx() *gorm.DB {
	return u.Begin()
}

func (u *roleRepo) AddRoleRepo(role *models.CreateMenuPermRequest) error {
	return u.DB.Model(models.Role{}).Create(&role.Role).Error
}

func (u *roleRepo) BatchCreateMenuPermRepo(mps *[]*models.RoleMenuPermission) error {
	return u.DB.Model(models.Role{}).Create(mps).Error
}

func (u *roleRepo) UpdateRoleRepo(role *models.Role) error {
	return u.DB.Model(models.Role{}).Updates(role).Error
}

func (u *roleRepo) DeleteRoleRepo(role *models.Role) error {
	return u.DB.Model(models.Role{}).Delete(role).Error
}

func (u *roleRepo) ListRoleRepo(page, perPage, userId int) (*[]models.Role, error) {
	var roles []models.Role
	return &roles, u.DB.Model(models.Role{}).
		Where("delete_at is null and created_user_id = ?", userId).
		Scan(&roles).Error
}

func buildCreateMenuPermRequest(r *[]models.MenuPermResponse) *models.CreateMenuPermRequest {
	var (
		MPermMap = make(map[int]*models.CreateMenuPermRequest)
		resp     models.CreateMenuPermRequest
	)
	for i := range *r {
		ele := (*r)[i]
		var rmp models.RoleMenuPermission
		rmp.RoleID = ele.RoleID
		rmp.MenuID = ele.MenuID
		rmp.PermissionID = ele.PermissionID

		if resp.ID != 0 {
			resp.MenuPerms = append(resp.MenuPerms, &rmp)
		} else {
			menuPerms := []*models.RoleMenuPermission{&rmp}
			resp = models.CreateMenuPermRequest{Role: ele.Role, MenuPerms: menuPerms}
			MPermMap[ele.ID] = &resp
		}
	}
	return &resp
}

func (u *roleRepo) RoleDetailRepo(roleId, userId int) (*models.CreateMenuPermRequest, error) {
	var roles []models.MenuPermResponse
	err := u.DB.Model(models.Role{}).
		Select("role.*, role_menu_permission.menu_id, role_menu_permission.role_id, role_menu_permission.permission_id").
		Joins("left join role_menu_permission on role_menu_permission.role_id = role.id").
		Where("role_menu_permission.delete_at is null and role.id = ?", roleId).
		Scan(&roles).Error
	if err != nil {
		return nil, err
	}
	return buildCreateMenuPermRequest(&roles), err
}

func (u *roleRepo) DeleteMenuPermissionByRoleIDRepo(roleId int) error {
	rmp := new(models.RoleMenuPermission)
	return u.DB.Model(models.RoleMenuPermission{}).
		Where("role_id = ? and deleted_at is null ", roleId).
		Delete(rmp).Error
}
