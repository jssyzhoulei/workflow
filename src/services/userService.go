package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/repositories"
)

type UserServiceI interface {
	AddUserSvc(ctx context.Context, user models.User) (models.User, error)
}

type userService struct {
	userRepo repositories.UserRepoI
}

func NewUserService(repos repositories.RepoI) UserServiceI {
	return &userService{
		userRepo: repos.GetUserRepo(),
	}
}

func (u *userService) AddUserSvc(ctx context.Context, user models.User) (models.User, error) {
	return user, u.userRepo.AddUserRepo(user)
}
