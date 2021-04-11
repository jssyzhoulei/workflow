package repositories

import (
	"errors"
	"fmt"
	"github.com/jssyzhoulei/workflow/src/models"
	"gorm.io/gorm"
	"math"
	"strings"
	"time"
)

//type UserRepoInterface interface {
//	AddUserRepo(user models.User, tx *gorm.DB) (int, error)
//	GetUserByIDRepo(id int) (user models.User, err error)
//	UpdateUserByIDRepo(user models.User, tx *gorm.DB) error
//	DeleteUserByIDRepo(id int, tx *gorm.DB) error
//	AddUserRoleRepo(userRole models.UserRole) error
//	GetUserListRepo(user models.User, page *models.Page, tx *gorm.DB, groupIds ...int64) ([]models.User, error)
//	BatchDeleteUsersRepo(ids []int64, tx *gorm.DB) error
//	GetTx() *gorm.DB
//	GetUsersByLoginNames([]string) ([]models.User, error)
//	AddUsersRepo(users []models.User, tx *gorm.DB) ([]int, error)
//	AddUserRolesRepo(roles []models.UserRole, tx *gorm.DB) error
//	DeleteUserRolesByUserId(ids []int, tx *gorm.DB) error
//	DeleteUserRolesById(id int, tx *gorm.DB) error
//	DeleteUserRolesByUserIds(ids []int64, tx *gorm.DB) error
//	UpdateUserRolesRepo(userRolesDTO models.UserRolesDTO, tx *gorm.DB) error
//	ImportUsersByGroupIdRepo(groupId int, userId []int) error
//	GetRoleIdsById(id int) ([]int, error)
//	GetRoleIdsByUserIds(ids []int) ([]int, error)
//	GetUsersRepo(condition *models.UserQueryByCondition) ([]*models.UserListResult, int64, error)
//	DeleteUserRolesRepo(userRolesDTO models.UserRolesDTO, tx *gorm.DB) error
//}

func (w *WorkRepo) GetTx() *gorm.DB {
	return w.Begin()
}

// AddUserRepo 添加用户
func (w *WorkRepo) AddUserRepo(user models.User, tx *gorm.DB) (int, error) {
	//userRecord, err := u.GetUserByName(user.LoginName)
	//if err != nil && userRecord.ID == 0 {
	//	return u.Create(&user).Error
	//}
	//return errors.New("user is exist")
	var db *gorm.DB
	if tx == nil {
		db = w.DB
	} else {
		db = tx
	}

	_, err := w.GetUserByName(user.LoginName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.ID, db.Create(&user).Error
		} else {
			return 0, err
		}
	} else {
		return 0, errors.New("user already exist")
	}
}

// GetUserByIDRepo 通过ID获取用户详情
func (w *WorkRepo) GetUserByIDRepo(id int) (user models.User, err error) {
	err = w.First(&user, id).Error
	return
}

func (w *WorkRepo) GetRoleIdsById(id int) ([]int, error) {
	var userRoles []models.UserRole
	var roleIds []int
	err := w.Model(&models.UserRole{}).Where("user_id=?", id).Find(&userRoles)
	if err != nil {
		return roleIds, err.Error
	}
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleID)
	}
	return roleIds, nil
}

// UpdateUserByIDRepo 根据ID编辑用户
func (w *WorkRepo) UpdateUserByIDRepo(user models.User, tx *gorm.DB) error {
	var (
		db = w.DB
	)
	if tx != nil {
		db = tx
	}
	userRecord, err := w.GetUserByName(user.LoginName)
	if err != nil || userRecord.ID == user.ID {
		return db.Model(&user).Updates(user).Error
	}
	return errors.New("user is exist")
}

func (w *WorkRepo) ImportUsersByGroupIdRepo(groupId int, userId []int) error {

	err := w.Model(&models.User{}).Where("id in ?", userId).Update("group_id", groupId)
	if err != nil {
		return err.Error
	}
	return nil
}

// DeleteUserByIDRepo 根据ID删除用户
func (w *WorkRepo) DeleteUserByIDRepo(id int, tx *gorm.DB) error {
	//var(
	//	user models.User
	//)
	var db *gorm.DB
	if tx == nil {
		db = w.DB
	} else {
		db = tx
	}
	updateColumnMap := map[string]interface{}{
		"status":     1,
		"deleted_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	err := db.Model(&models.User{}).Where("id=?", id).Updates(updateColumnMap).Error
	if err != nil {
		return err
	}
	return nil

	//if id != 0 {
	//	user.ID = id
	//	return db.Delete(&user).Error
	//}
	//return nil
}

// GetUserListRepo 获取用户列表
func (w *WorkRepo) GetUsersRepo(condition *models.UserQueryByCondition) ([]*models.UserListResult, int64, error) {
	var err error
	db := w.DB

	whereCondition := " where 1=1 and a.deleted_at IS NULL and b.deleted_at IS NULL "
	var conditionVal = make(map[string]interface{})

	if len(condition.ID) != 0 {
		whereCondition += " and a.id in @ids"
		conditionVal["ids"] = condition.ID
	}
	if len(condition.GroupId) != 0 {
		whereCondition += " and group_id in @group_ids"
		conditionVal["group_ids"] = condition.GroupId
	}
	if len(condition.LoginName) != 0 {
		whereCondition += " and login_name like @login_names"
		conditionVal["login_names"] = "%" + condition.LoginName + "%"
	}

	orderSql := " order by a.id desc "

	page := condition.PageNum
	limit := condition.PageSize
	offset := page*limit - limit

	var pageSql string
	if limit != 0 {
		pageSql = fmt.Sprintf(" limit %d offset %d ", limit, offset)
	}

	countSQl := "select count(1) as count from (%s) a"

	sqlStr := `
SELECT DISTINCT
	a.id,
	a.created_at,
	a.user_name,
	a.login_name,
	a.mobile,
	d.name AS group_name,
   GROUP_CONCAT(
   c.name
   ) as role_name,
	d.id AS group_id
FROM
	` + "`user`" + ` a
	LEFT JOIN user_role b ON a.id = b.user_id
	LEFT JOIN ` + "`role`" + ` c ON b.role_id = c.id and c.status = 0
	LEFT JOIN ` + "`group`" + ` d ON a.group_id = d.id
`
	fullSql := sqlStr + whereCondition + " group by a.id " + orderSql
	totalSql := fmt.Sprintf(countSQl, fullSql)

	var resultScan = make([]*models.UserListScanResult, 0)
	err = db.Raw(fullSql+pageSql, conditionVal).Scan(&resultScan).Error
	if err != nil {
		return nil, 0, err
	}
	var total int64
	err = db.Raw(totalSql, conditionVal).Scan(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 处理结果
	var result = make([]*models.UserListResult, 0)
	for _, val := range resultScan {
		_tmp := &models.UserListResult{
			Id:        val.Id,
			LoginName: val.LoginName,
			CreatedAt: val.CreatedAt,
			UserName:  val.UserName,
			GroupName: val.GroupName,
			RoleName:  strings.Split(val.RoleName, ","),
			GroupId:   val.GroupId,
			Mobile:    val.Mobile,
		}
		result = append(result, _tmp)
	}

	return result, total, nil
}

// GetUserListRepo 获取用户列表
func (w *WorkRepo) GetUserListRepo(user models.User, page *models.Page, tx *gorm.DB, groupIds ...int64) ([]models.User, error) {
	var (
		users []models.User
	)
	var err error
	var db *gorm.DB
	if tx == nil {
		db = w.DB
	} else {
		db = tx
	}

	db = db.Model(&models.User{}).
		Select("user_name, group_id, created_at, id, login_name, mobile, user_type, status")

	if user.UserName != "" {
		db = db.Where("user_name like ?", "%"+user.UserName+"%")
	}

	if user.ID != 0 {
		db = db.Where("id=?", user.ID)
	}
	if user.GroupID != 0 {
		db = db.Where("group_id = ?", user.GroupID)
	} else if len(groupIds) > 0 {
		db = db.Where("group_id in ?", groupIds)
	}
	if page != nil {
		//db.DB()
		if page.PageNum == 0 {
			page.PageNum = 1
		}
		if page.PageSize == 0 {
			page.PageSize = 10
		}
		err := db.Count(&page.Total).Error
		if err != nil {
			return nil, err
		}
		page.TotalPage = math.Ceil(float64(page.Total)) / float64(page.PageSize)
		err = db.Order("id desc").Limit(page.PageSize).Offset(page.PageSize*page.PageNum - page.PageSize).Find(&users).Error
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

func (w *WorkRepo) GetRoleIdsByUserIds(ids []int) ([]int, error) {
	var userRoles []models.UserRole
	var roleIds []int
	err := w.Model(&models.UserRole{}).Where("user_id in ?", ids).Find(&userRoles).Error
	if err != nil {
		return roleIds, err
	}
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleID)
	}
	return roleIds, nil
}

// AddUserRoleRepo ...
func (w *WorkRepo) AddUserRoleRepo(userRole models.UserRole) error {
	return w.Create(&userRole).Error
}

// BatchDeleteUsersRepo 批量删除用户
func (w *WorkRepo) BatchDeleteUsersRepo(ids []int64, tx *gorm.DB) error {
	var db *gorm.DB
	if tx == nil {
		db = w.DB
	} else {
		db = tx
	}
	return db.Model(&models.User{}).Where("id in ?", ids).Delete(&models.User{}).Error
}

// GetUserByName 根据用户名获取用户
func (w *WorkRepo) GetUserByName(name string) (*models.User, error) {
	var (
		user = new(models.User)
		err  error
	)
	err = w.Where("login_name=?", name).First(&user).Error
	return user, err
}

func (w *WorkRepo) GetUsersByLoginNames(loginNames []string) ([]models.User, error) {
	var (
		users []models.User
		err   error
	)
	err = w.Table("user").Select("*").Where("login_name in ?", loginNames).Find(&users).Error
	return users, err
}

func (w *WorkRepo) AddUsersRepo(users []models.User, tx *gorm.DB) ([]int, error) {
	var (
		db  = w.DB
		err error
		ids []int
	)
	if tx != nil {
		db = tx
	}
	if users != nil {
		err = db.Create(&users).Error
	}

	if err == nil {
		for _, user := range users {
			ids = append(ids, user.ID)
		}
	}
	return ids, err
}

func (w *WorkRepo) AddUserRolesRepo(roles []models.UserRole, tx *gorm.DB) error {
	var (
		db  = w.DB
		err error
	)
	if tx != nil {
		db = tx
	}
	if roles != nil {
		err = db.Create(&roles).Error
	}
	return err
}

func (w *WorkRepo) DeleteUserRolesByUserId(ids []int, tx *gorm.DB) error {
	var (
		db = w.DB
	)
	if tx != nil {
		db = tx
	}
	return db.Table("user_role").Where("user_id IN ?", ids).Where("deleted_at is NULL").Delete(&models.UserRole{}).Error
}

func (w *WorkRepo) DeleteUserRolesById(id int, tx *gorm.DB) error {
	var (
		db = w.DB
	)
	if tx != nil {
		db = tx
	}
	return db.Table("user_role").Where("user_id = ?", id).Where("deleted_at is NULL").Delete(&models.UserRole{}).Error
}

func (w *WorkRepo) DeleteUserRolesByUserIds(ids []int64, tx *gorm.DB) error {
	var (
		db = w.DB
	)
	if tx != nil {
		db = tx
	}
	return db.Table("user_role").Where("user_id IN ?", ids).Where("deleted_at is NULL").Delete(&models.UserRole{}).Error
}

func (w *WorkRepo) UpdateUserRolesRepo(userRolesDTO models.UserRolesDTO, tx *gorm.DB) error {
	var (
		db  = w.DB
		err error
	)
	if tx != nil {
		db = tx
	}

	for _, roleId := range userRolesDTO.RoleIDs {
		db.Table("user_role").Where("user_id = ?", userRolesDTO.ID).Updates(map[string]interface{}{"user_id": userRolesDTO.ID, "role_id": roleId})
	}

	return err
}

func (w *WorkRepo) DeleteUserRolesRepo(userRolesDTO models.UserRolesDTO, tx *gorm.DB) error {
	var (
		db  = w.DB
		err error
	)
	if tx != nil {
		db = tx
	}
	err = db.Table("user_role").Where("user_id = ?", userRolesDTO.ID).Delete(&models.UserRole{}).Error
	if err != nil {
		return err
	}
	return nil
}
