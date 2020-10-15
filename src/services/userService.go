package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/md5"
)

// UserServiceI 用户服务接口
type UserServiceInterface interface {
	AddUserSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error)
	GetUserByIDSvc(ctx context.Context, id int) (models.User, error)
	UpdateUserByIDSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error)
	DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error)
	GetUserListSvc(ctx context.Context, user models.User, page *models.Page ) ([]models.User, error)
}

// UserService 用户服务，实现 UserServiceInterface
type userService struct {
	userRepo repositories.UserRepoInterface
}

// NewUserService UserService 构造函数
func NewUserService(repos repositories.RepoI) UserServiceInterface {
	return &userService{
		userRepo: repos.GetUserRepo(),
	}
}

// AddUserSvc 添加用户
func (u *userService) AddUserSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	user.Password = md5.EncodeMD5(user.Password)
	err = u.userRepo.AddUserRepo(user)
	return pb_user_v1.NullResponse{}, err
}

// GetUserByIDSvc 获取用户详情
func (u *userService) GetUserByIDSvc(ctx context.Context, id int) (models.User, error) {
	var (
		user models.User
		err error
	)
	user, err = u.userRepo.GetUserByIDRepo(id)
	return user, err
}

// UpdateUserByIDSvc 根据ID编辑用户
func (u *userService) UpdateUserByIDSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error) {
	err := u.userRepo.UpdateUserByIDRepo(user)
	return pb_user_v1.NullResponse{}, err
}

// DeleteUserByID 根据ID删除用户信息
func (u *userService) DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = u.userRepo.DeleteUserByIDRepo(id)
	return pb_user_v1.NullResponse{}, err
}

// GetUserListSvc 获取用户列表
func (u *userService) GetUserListSvc(ctx context.Context, user models.User, page *models.Page ) ([]models.User, error){
	var(
		users []models.User
		err error
	)
	users, err = u.userRepo.GetUserListRepo(user, page)
	if err != nil {
		return nil, err
	}
	return users, nil
}
