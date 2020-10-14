package repositories

import (
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/md5"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
)

type UserRepoI interface {
	AddUserRepo(user models.User) error
}

type userRepo struct {
	repo *yorm.DBRepo
	*gorm.DB
}

func NewUserRepo(db *yorm.DB) UserRepoI {
	return &userRepo{
		repo: db.GetRepo("user"),
		DB: db.DB,
	}
}

func (u *userRepo) AddUserRepo(user models.User) error {
	var err error
	var db = u.DB
	var count int64
	err = db.Model(&models.User{}).Where("user_name=?", user.UserName).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("创建用户：%s 失败，已经存在", user.UserName)
	}

	user.Password = md5.EncodeMD5(user.Password)
	db = db.Create(&user)

	return u.repo.AddQuery("user", user).Exec("addUser").Err()
}

