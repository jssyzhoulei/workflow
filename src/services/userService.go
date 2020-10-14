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
	AddUserSvc(ctx context.Context, user *pb_user_v1.UserProto) (*pb_user_v1.NullResponse, error)
	GetUserByIDSvc(ctx context.Context, id int) (models.User, error)
	UpdateUserByIDSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error)
	DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error)
}

// UserService 用户服务，实现 UserServiceInterface
type UserService struct {
	userRepo repositories.UserRepoInterface
}

// NewUserService UserService 构造函数
func NewUserService(repos repositories.RepoI) UserServiceInterface {
	return &UserService{
		userRepo: repos.GetUserRepo(),
	}
}

// AddUserSvc 添加用户
func (u *UserService) AddUserSvc(ctx context.Context, user *pb_user_v1.UserProto) (*pb_user_v1.NullResponse, error) {
	var err error
	tx := u.userRepo.GetTx()
	defer func() {
		if r := recover(); r != nil{
			tx.Rollback()
		}
	}()

	newUser := &models.User{
		UserName: user.UserName,
		Password: md5.EncodeMD5(user.Password),
		LoginName: user.LoginName,
		Mobile: int(user.Mobile),
	}

	err = u.userRepo.AddUserRepo(newUser, tx)
	if err != nil {
		return &pb_user_v1.NullResponse{Code: 1}, err
	}
	tx.Commit()
	return &pb_user_v1.NullResponse{Code : 0}, nil
}

// GetUserByIDSvc 获取用户详情
func (u *UserService) GetUserByIDSvc(ctx context.Context, id int) (models.User, error) {
	var (
		user models.User
		err error
	)
	user, err = u.userRepo.GetUserByIDRepo(id)
	return user, err
}

// UpdateUserByIDSvc 根据ID编辑用户
func (u *UserService) UpdateUserByIDSvc(ctx context.Context, user models.User) (pb_user_v1.NullResponse, error) {
	err := u.userRepo.UpdateUserByIDRepo(user)
	return pb_user_v1.NullResponse{}, err
}

// DeleteUserByID 根据ID删除用户信息
func (u *UserService) DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	err = u.userRepo.DeleteUserByIDRepo(id)
	return pb_user_v1.NullResponse{}, err
}
