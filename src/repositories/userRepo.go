package repositories

import (
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
	"math"
	"time"
)

// UserRepoI ...
type UserRepoInterface interface {
	AddUserRepo(user models.User, tx *gorm.DB) (int, error)
	GetUserByIDRepo(id int) (user models.User, err error)
	UpdateUserByIDRepo(user models.User, tx *gorm.DB) error
	DeleteUserByIDRepo(id int,tx *gorm.DB) error
	AddUserRoleRepo(userRole models.UserRole) error
	GetUserListRepo(user models.User, page *models.Page, tx *gorm.DB, groupIds ...int64) ([]models.User, error)
	BatchDeleteUsersRepo(ids []int64, tx *gorm.DB) error
	GetTx() *gorm.DB
	GetUsersByLoginNames([]string) ([]models.User, error)
	AddUsersRepo(users []models.User, tx *gorm.DB) ([]int, error)
	AddUserRolesRepo(roles []models.UserRole, tx *gorm.DB) error
	DeleteUserRolesByUserId(ids []int, tx *gorm.DB) error
	DeleteUserRolesById(id int, tx *gorm.DB) error
	DeleteUserRolesByUserIds(ids []int64, tx *gorm.DB) error
	UpdateUserRolesRepo (userRolesDTO models.UserRolesDTO, tx *gorm.DB) error
}

type userRepo struct {
	*gorm.DB
}

func (u *userRepo) GetTx() *gorm.DB {
	return u.Begin()
}

// NewUserRepo ...
func NewUserRepo(db *yorm.DB) UserRepoInterface {
	return &userRepo{
		DB: db.DB,
	}
}

// AddUserRepo 添加用户
func (u *userRepo) AddUserRepo(user models.User, tx *gorm.DB) (int, error) {
	//userRecord, err := u.GetUserByName(user.LoginName)
	//if err != nil && userRecord.ID == 0 {
	//	return u.Create(&user).Error
	//}
	//return errors.New("user is exist")
	var db *gorm.DB
	if tx == nil {
		db = u.DB
	} else {
		db = tx
	}

	_, err := u.GetUserByName(user.LoginName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(&user).Error
			if err != nil {
				fmt.Println("create user error: ", err.Error())
			}
			return user.ID, nil
		} else {
			fmt.Println(err.Error(), "=================")
		}
	}
	return 0, err
}

// GetUserByIDRepo 通过ID获取用户详情
func (u *userRepo) GetUserByIDRepo(id int) (user models.User, err error) {
	err = u.First(&user, id).Error
	return
}

// UpdateUserByIDRepo 根据ID编辑用户
func (u *userRepo) UpdateUserByIDRepo(user models.User, tx *gorm.DB) error {
	var (
		db = u.DB
	)
	if tx != nil {
		db = tx
	}
	userRecord, err := u.GetUserByName(user.UserName)
	if err != nil || userRecord.ID == user.ID {
		return db.Model(&user).Updates(user).Error
	}
	return errors.New("user is exist")
}

// DeleteUserByIDRepo 根据ID删除用户
func (u *userRepo) DeleteUserByIDRepo(id int,tx *gorm.DB) error {
	//var(
	//	user models.User
	//)
	var db *gorm.DB
	if tx == nil {
		db = u.DB
	} else {
		db = tx
	}
	updateColumnMap := map[string]interface{} {
		"status": 1,
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
func (u *userRepo) GetUserListRepo(user models.User, page *models.Page, tx *gorm.DB, groupIds ...int64) ([]models.User, error){
	var(
		users []models.User
	)
	var err error
	var db *gorm.DB
	if tx == nil {
		db = u.DB
	} else {
		db = tx
	}
	dbPage := *u.DB
	db = u.Table("user").
		Select("user_name, group_id, created_at, id, login_name, mobile, user_type")

	if user.UserName != "" {
		db = db.Where("user_name like ?", "%" + user.UserName + "%")
	}

	if user.ID != 0 {
		db = db.Where("id=?", user.ID)
	}
	if user.GroupID != 0 {
		db = db.Where("group_id = ?", user.GroupID)
	}else if len(groupIds) > 0 {
		db = db.Where("group_id in ?", groupIds)
	}
	if page != nil {
		db.DB()
		if page.PageNum == 0 {
			page.PageNum = 1
		}
		if page.PageSize == 0 {
			page.PageSize = 10
		}
		err := dbPage.Table("(?) as p",db).Count(&page.Total).Error
		if err != nil {
			return nil, err
		}
		page.TotalPage = math.Ceil(float64(page.Total)) / float64(page.PageSize)
		err = db.Limit(page.PageSize).Offset(page.PageSize * (page.PageNum - 1)).Find(&users).Error
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

// AddUserRoleRepo ...
func (u *userRepo) AddUserRoleRepo(userRole models.UserRole) error {
	return u.Create(&userRole).Error
}

// BatchDeleteUsersRepo 批量删除用户
func (u *userRepo) BatchDeleteUsersRepo(ids []int64, tx *gorm.DB) error {
	var db *gorm.DB
	if tx == nil {
		db = u.DB
	} else {
		db = tx
	}
	return db.Model(&models.User{}).Where("id in ?", ids).Delete(&models.User{}).Error
}
// GetUserByName 根据用户名获取用户
func (u *userRepo) GetUserByName(name string)(models.User, error) {
	var(
		user = new(models.User)
		err error
	)
	err = u.Where("login_name=?", name).First(&user).Error
	return *user, err
}

func (u *userRepo) GetUsersByLoginNames(loginNames []string) ([]models.User, error) {
	var (
		users []models.User
		err error
	)
	err = u.Table("user").Select("*").Where("login_name In ?", loginNames).Find(&users).Error
	return users, err
}

func (u *userRepo) AddUsersRepo(users []models.User, tx *gorm.DB) ([]int, error) {
	var (
		db = u.DB
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

func (u *userRepo) AddUserRolesRepo(roles []models.UserRole, tx *gorm.DB) error {
	var (
		db = u.DB
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

func (u *userRepo) DeleteUserRolesByUserId(ids []int, tx *gorm.DB) error {
	var (
		db = u.DB
	)
	if tx != nil {
		db = tx
	}
	return db.Table("user_role").Where("user_id IN ?", ids).Where("deleted_at is NULL").Delete(&models.UserRole{}).Error
}

func (u *userRepo) DeleteUserRolesById(id int, tx *gorm.DB) error {
	var (
		db = u.DB
	)
	if tx != nil {
		db = tx
	}
	return db.Table("user_role").Where("user_id = ?", id).Where("deleted_at is NULL").Delete(&models.UserRole{}).Error
}

func (u *userRepo) DeleteUserRolesByUserIds(ids []int64, tx *gorm.DB) error {
	var (
		db = u.DB
	)
	if tx != nil {
		db = tx
	}
	return db.Table("user_role").Where("user_id IN ?", ids).Where("deleted_at is NULL").Delete(&models.UserRole{}).Error
}

func (u *userRepo) UpdateUserRolesRepo (userRolesDTO models.UserRolesDTO, tx *gorm.DB) error {
	var(
		db = u.DB
		err error
	)
	if tx != nil {
		db = tx
	}

	for _, roleId := range userRolesDTO.RoleIDs  {
		db.Table("user_role").Where("user_id = ?", userRolesDTO.ID).Updates(map[string]interface{}{"user_id": userRolesDTO.ID, "role_id": roleId})
	}

	return err
}