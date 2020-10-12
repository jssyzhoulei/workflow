package repositories

import (
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
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
	fmt.Println("我到了")
	fmt.Println(user)
	return u.repo.AddQuery("user", user).Exec("addUser").Err()
}

