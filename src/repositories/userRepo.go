package repositories

import (
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
)

// UserRepoI ...
type UserRepoInterface interface {
	AddUserRepo(user *models.User, tx *gorm.DB) error
	GetUserByIDRepo(id int) (user models.User, err error)
	UpdateUserByIDRepo(user models.User) error
	DeleteUserByIDRepo(id int) error
	GetTx() *gorm.DB
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
func (u *userRepo) AddUserRepo(user *models.User, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = u.DB
	} else {
		db = tx
	}

	//判断该用户是否存在
	var count int64
	err = db.Model(&models.User{}).Where("user_name=?", user.UserName).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("创建用户：%s 失败，已经存在", user.UserName)
	}

	//创建新的用户
	newUserInfo := &models.User{
		UserName: user.UserName,
		Password: user.Password,
		LoginName: user.LoginName,
		GroupID: user.GroupID,
		Mobile: user.Mobile,
		UserType: user.UserType,
	}
	if err = db.Create(newUserInfo).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByIDRepo 获取用户详情
func (u *userRepo) GetUserByIDRepo(id int) (user models.User, err error) {
	err = u.First(&user, id).Error
	return
}

// UpdateUserByIDRepo 根据ID编辑用户
func (u *userRepo) UpdateUserByIDRepo(user models.User) error {
	return u.Model(&user).Updates(user).Error
}

// DeleteUserByIDRepo 根据ID删除用户
func (u *userRepo) DeleteUserByIDRepo(id int) error {
	var (
		user models.User
	)
	if id != 0 {
		user.ID = id
		return u.Delete(&user).Error
	}
	return nil
}