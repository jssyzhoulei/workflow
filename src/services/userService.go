package services

import (
	"context"
	"fmt"
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
	GetUserListSvc(ctx context.Context, user *pb_user_v1.UserPage) (c *pb_user_v1.UsersPage, err error)
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
	var err error
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
	fmt.Printf("%+v",user)
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
func (u *userService) GetUserListSvc(ctx context.Context, userPage *pb_user_v1.UserPage) (c *pb_user_v1.UsersPage, err error){
	var (
		page models.Page
		user models.User
	)
	if userPage.Page != nil {
		page.PageSize = int(userPage.Page.PageSize)
		page.PageNum = int(userPage.Page.PageNum)
	}
	if userPage.User != nil {
		user.UserName = userPage.User.UserName
	}
	users, err := u.userRepo.GetUserListRepo(user, &page, nil)
	if err != nil {
		return c, err
	}
	c = &pb_user_v1.UsersPage{}
	c.Users = &pb_user_v1.Users{}
	for _, user := range users {
		var userProto pb_user_v1.UserProto
		userProto.Id = &pb_user_v1.Index{
			Id:                   int64(user.ID),
		}

		c.Users.Users = append(c.Users.Users, &userProto)
	}
	return c, nil
}
