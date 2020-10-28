package repositories

import (
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
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
	ListRolesRepo(pageObj *pb_user_v1.RolePageRequestProto, userId int) (*pb_user_v1.RolePageRequestProto, error)
	BuildPermissionTree(module int) ([]*pb_user_v1.Cascade, error)
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
	return u.DB.Model(models.RoleMenuPermission{}).Create(mps).Error
}

func (u *roleRepo) UpdateRoleRepo(role *models.Role) error {
	return u.DB.Model(models.Role{}).Where("id = ?", role.ID).Updates(role).Error
}

func (u *roleRepo) DeleteRoleRepo(role *models.Role) error {
	var count int64
	u.Model(models.UserRole{}).Where("role_id = ? and deleted_at is null ", role.ID).Count(&count)
	if count != 0{
		return errors.New("relation user")
	}
	return u.DB.Model(models.Role{}).Delete(role).Error
}

func (u *roleRepo) ListRoleRepo(page, perPage, userId int) (*[]models.Role, error) {
	var roles []models.Role
	return &roles, u.DB.Model(models.Role{}).
		Where("deleted_at is null and created_user_id = ?", userId).
		Scan(&roles).Error
}

type roleUserIds struct {
	models.Role
	UserID int    `json:"user_id"`
	IDs    string `json:"ids"`
}

func buildRoleProto(roles *[]roleUserIds) *[]*pb_user_v1.RoleProto {

	var rolePbs []*pb_user_v1.RoleProto
	for _, r := range *roles {
		one := pb_user_v1.RoleProto{
			Name:       r.Name,
			Remark:     r.Remark,
			DataPermit: int32(r.DataPermit),
			Status:     int32(r.Status),
			Id:         int64(r.ID),
			CreatedAt:  r.CreatedAt.Format("2006-01-02 15:04:05"),
			Ids:        r.IDs,
		}
		rolePbs = append(rolePbs, &one)
	}

	return &rolePbs
}

func (u *roleRepo) ListRolesRepo(pageObj *pb_user_v1.RolePageRequestProto, userId int) (*pb_user_v1.RolePageRequestProto, error) {
	var (
		page  int = 1
		limit int = 10
		name  string
		roles []roleUserIds
		resp  pb_user_v1.RolePageRequestProto
	)

	resp.Page = new(pb_user_v1.Page)
	if pageObj != nil {
		if pageObj.Page.PageNum != 0 {
			page = int(pageObj.Page.PageNum)
			limit = int(pageObj.Page.PageSize)
		}
		name = pageObj.Name
	}

	u.DB.Raw(`SELECT count(1) count from role WHERE role.deleted_at is null
				and role.name like ? `, "%"+name+"%").
		Count(&resp.Page.Total)

	err := u.DB.Raw(`SELECT role.id, role.status, role.created_at, role.name, role.remark,
						role.data_permit, group_concat(DISTINCT(user_role.user_id)) ids FROM role 
						left join user_role on user_role.role_id = role.id and user_role.deleted_at is null 
						WHERE role.deleted_at is null 
						and role.name like ?  
						group by role.id 
						LIMIT ? OFFSET ?`, "%"+name+"%", limit, limit*(page-1)).
		Scan(&roles).Error

	if err == nil {
		resp.Roles = *buildRoleProto(&roles)
	}

	return &resp, err
}

func buildCreateMenuPermRequest(r *[]models.MenuPermResponse) *models.CreateMenuPermRequest {
	var resp models.CreateMenuPermRequest
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
		}
	}
	return &resp
}

func (u *roleRepo) RoleDetailRepo(roleId, userId int) (*models.CreateMenuPermRequest, error) {
	var roles []models.MenuPermResponse
	err := u.DB.Model(models.Role{}).
		Select("role.*, role_menu_permission.menu_id, role_menu_permission.role_id, role_menu_permission.permission_id").
		Joins("left join role_menu_permission on role_menu_permission.role_id = role.id").
		Where("role_menu_permission.deleted_at is null and role.id = ?", roleId).
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

type MenuPermissionTree struct {
	MenuID       int    `json:"menu_id"`
	ParentID     int    `json:"parent_id"`
	MenuName     string `json:"menu_name"`
	PermissionID int    `json:"permission_id"`
	UriName      string `json:"uri_name"`
}

func (u *roleRepo) BuildPermissionTree(module int) ([]*pb_user_v1.Cascade, error) {
	var tree []MenuPermissionTree
	u.Raw(`select m.id menu_id, m.name menu_name, m.parent_id, p.id permission_id, p.uri_name 
				from menu m left join permission p on p.menu_id = m.id 
				where m.deleted_at is null and p.deleted_at is null and m.module = ? `, module).Scan(&tree)
	return buildCas(&tree)
}

func buildCas(tree *[]MenuPermissionTree) ([]*pb_user_v1.Cascade, error) {
	// 顶层菜单
	var cas []*pb_user_v1.Cascade
	// menu id 和 Menu 的映射
	menuIdCasMap := make(map[int]*pb_user_v1.Cascade)
	// parent id 和 子Menu列表 的映射
	parentIdCasMap := make(map[int][]*pb_user_v1.Cascade)
	for _, i := range *tree {
		if ca, ok := menuIdCasMap[i.MenuID]; ok {
			// permission ca
			ca.Child = append(ca.Child, &pb_user_v1.Cascade{
				Label: i.UriName,
				Value: int64(i.PermissionID),
			})
		} else {
			// 这是menu ca
			var ca pb_user_v1.Cascade
			ca.Label = i.MenuName
			ca.Value = int64(i.MenuID)
			ca.Child = []*pb_user_v1.Cascade{
				&pb_user_v1.Cascade{
					Label: i.UriName,
					Value: int64(i.PermissionID),
				},
			}
			menuIdCasMap[i.MenuID] = &ca
			if i.ParentID == -1 {
				cas = append(cas, &ca)
			} else {
				if p, ok := parentIdCasMap[i.ParentID]; ok {
					p = append(p, &ca)
				} else {
					parentIdCasMap[i.ParentID] = []*pb_user_v1.Cascade{&ca}
				}
			}
		}
	}

	return cas, buildMenuChildren(cas, parentIdCasMap)
}

func buildMenuChildren(top []*pb_user_v1.Cascade, rel map[int][]*pb_user_v1.Cascade) error {
	for _, ca := range top {
		caL, ok := rel[int(ca.Value)]
		if !ok {
			continue
		}
		ca.Children = caL
		err := buildMenuChildren(caL, rel)
		if err != nil {
			return err
		}
	}
	return nil
}
