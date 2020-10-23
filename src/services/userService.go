package services

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/config"
)

// UserServiceI 用户服务接口
type UserServiceInterface interface {
	AddUserSvc(ctx context.Context, userRolesDTO models.UserRolesDTO) (pb_user_v1.NullResponse, error)
	GetUserByIDSvc(ctx context.Context, id int) (c *pb_user_v1.UserProto, err error)
	UpdateUserByIDSvc(ctx context.Context, userRolesDTO models.UserRolesDTO) (pb_user_v1.NullResponse, error)
	DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error)
	AddUsersSvc(ctx context.Context, users *pb_user_v1.AddUsersRequest) (pb_user_v1.NullResponse, error)
	GetUserListSvc(ctx context.Context, user *pb_user_v1.UserPage) (c *pb_user_v1.UsersPage, err error)
	BatchDeleteUsersSvc(ctx context.Context, ids []int64) (pb_user_v1.NullResponse, error)
	ImportUsersByGroupIdSvc(ctx context.Context, id models.GroupAndUserId) (pb_user_v1.NullResponse, error)
	GetUsersSvc(ctx context.Context, data *pb_user_v1.UserQueryCondition) (*pb_user_v1.UserQueryResponse, error)
}

// UserService 用户服务，实现 UserServiceInterface
type userService struct {
	userRepo repositories.UserRepoInterface
	config   *config.Config
	roleRepo repositories.RoleRepoI
}

// NewUserService UserService 构造函数
func NewUserService(repos repositories.RepoI, c *config.Config) UserServiceInterface {
	return &userService{
		userRepo: repos.GetUserRepo(),
		config: c,
		roleRepo: repos.GetRoleRepo(),
	}
}

// AddUserSvc 添加用户
func (u *userService) AddUserSvc(ctx context.Context, userRolesDTO models.UserRolesDTO) (pb_user_v1.NullResponse, error) {
	var (
		userRoles    []models.UserRole
	)

	tx := u.userRepo.GetTx()
	key, _ := u.config.GetString("passwordKey")
	userRolesDTO.Password,_ = userRolesDTO.EncodePwd(key)
	id, err := u.userRepo.AddUserRepo(userRolesDTO.User, tx)
	if err != nil {
		tx.Rollback()
		return pb_user_v1.NullResponse{}, err
	}
	if id == 0 {
		tx.Commit()
		return pb_user_v1.NullResponse{}, nil
	}
	for _, roleId := range userRolesDTO.RoleIDs {
		userRoles = append(userRoles, models.UserRole{
			UserID: id,
			RoleID: int(roleId),
		})
	}

	err = u.userRepo.AddUserRolesRepo(userRoles, tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()




	//var err error
	//var userRole models.UserRole
	//key, _ := u.config.GetString("passwordKey")
	//user.Password,_ = user.EncodePwd(key)
	//id, err = u.userRepo.AddUserRepo(user)
	//
	//userRole.UserID = user.ID
	//err = u.userRepo.AddUserRoleRepo(userRole)
	return pb_user_v1.NullResponse{}, err
}

// GetUserByIDSvc 获取用户详情
func (u *userService) GetUserByIDSvc(ctx context.Context, id int) (c *pb_user_v1.UserProto, err error) {
	//var (
	//	user models.User
	//	err error
	//)
	//user, err = u.userRepo.GetUserByIDRepo(id)
	//return user, err
	var user models.User
	var userProto pb_user_v1.UserProto
	user, err = u.userRepo.GetUserByIDRepo(id)
	if err != nil {
		return c, err
	}
	roleIds, err := u.userRepo.GetRoleIdsById(id)
	if err != nil {
		return c, err
	}
	c = &pb_user_v1.UserProto{}
	userProto.Id = &pb_user_v1.Index{
		Id:          int64(user.ID),
	}
	userProto = pb_user_v1.UserProto{
		UserName: user.UserName,
		LoginName: user.LoginName,
		Password: user.Password,
		GroupId: int64(user.GroupID),
		UserType: int64(user.UserType),
		Ststus: int64(user.Status),
		Mobile: int64(user.Mobile),
	}
	for _, roleId := range roleIds {
		userProto.RoleIds = append(userProto.RoleIds, &pb_user_v1.Index{Id: int64(roleId)})
	}
	return c, nil
}

// UpdateUserByIDSvc 根据ID编辑用户
func (u *userService) UpdateUserByIDSvc(ctx context.Context, userRolesDTO models.UserRolesDTO) (pb_user_v1.NullResponse, error) {
	//var (
	//	userRoles    []models.UserRole
	//)

	tx := u.userRepo.GetTx()
	err := u.userRepo.UpdateUserByIDRepo(userRolesDTO.User, tx)
	if err != nil {
		tx.Rollback()
		return pb_user_v1.NullResponse{}, err
	}
	if userRolesDTO.ID == 0 {
		tx.Commit()
		return pb_user_v1.NullResponse{}, nil
	}
	//for _, roleId := range userRolesDTO.RoleIDs {
	//	userRoles = append(userRoles, models.UserRole{
	//		UserID: userRolesDTO.ID,
	//		RoleID: int(roleId),
	//	})
	//}

	err = u.userRepo.UpdateUserRolesRepo(userRolesDTO, tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return pb_user_v1.NullResponse{}, err
}

// ImportUsersByGroupIdSvc ...
func (u *userService) ImportUsersByGroupIdSvc(ctx context.Context, id models.GroupAndUserId) (pb_user_v1.NullResponse, error) {
	err := u.userRepo.ImportUsersByGroupIdRepo(id.GroupId, id.UserIds)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	return pb_user_v1.NullResponse{}, nil
}
// DeleteUserByID 根据ID删除用户信息
func (u *userService) DeleteUserByIDSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	tx := u.userRepo.GetTx()
	err = u.userRepo.DeleteUserRolesById(id, tx)
	if err != nil {
		tx.Rollback()
		return pb_user_v1.NullResponse{}, err
	}

	err = u.userRepo.DeleteUserByIDRepo(id, tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return pb_user_v1.NullResponse{}, err
}

// GetUserListSvc 获取用户列表

func (u *userService) GetUserListSvc(ctx context.Context, userPage *pb_user_v1.UserPage) (c *pb_user_v1.UsersPage, err error){
	var (
		page models.Page
		user models.User
		userIds []int
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
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}
	reoleIds, err := u.userRepo.GetRoleIdsByUserIds(userIds)
	if err != nil {
		return nil, err
	}
	c = &pb_user_v1.UsersPage{}
	c.Users = &pb_user_v1.Users{}
	for _, user := range users {
		var userProto pb_user_v1.UserProto
		//userProto.Id = &pb_user_v1.Index{
		//	Id:                   int64(user.ID),
		//}
		userProto = pb_user_v1.UserProto{
			Id: &pb_user_v1.Index{Id: int64(user.ID)},
			UserName: user.UserName,
			LoginName: user.LoginName,
			Password: user.Password,
			Mobile: int64(user.Mobile),
			GroupId: int64(user.GroupID),
			Ststus: int64(user.Status),
			UserType: int64(user.UserType),
		}

		for _, v := range reoleIds {
			userProto.RoleIds = append(userProto.RoleIds, &pb_user_v1.Index{Id: int64(v)})
		}
		c.Users.Users = append(c.Users.Users, &userProto)
	}
	return c, nil
}

// BatchDeleteUsersRepo ...
func (u *userService) BatchDeleteUsersSvc(ctx context.Context, ids []int64) (pb_user_v1.NullResponse, error) {
	var (
		err error
	)
	tx := u.userRepo.GetTx()
	err = u.userRepo.DeleteUserRolesByUserIds(ids, tx)
	if err != nil {
		tx.Rollback()
		return pb_user_v1.NullResponse{}, err
	}

	err = u.userRepo.BatchDeleteUsersRepo(ids, tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return pb_user_v1.NullResponse{}, err
}

// AddUsersSvc ...
func (u *userService) AddUsersSvc(ctx context.Context, usersReq *pb_user_v1.AddUsersRequest) (pb_user_v1.NullResponse, error) {
	//查找所有用户，如果用户是别的用户组下将不进行操作返回那个用户重复
	//如果是覆盖写将对已经存在的用户修改
	//如果是追加则只新增对应的用户
	//覆盖写将对用户角色进行逻辑删除后在插入新的角色
	var (
		names        []string
		users        []models.User
		userIsExist  []models.User
		userIdsIsExist []int
		userNotExist []models.User
		roleIds      []int64
		userRoles    []models.UserRole
	)
	if usersReq.Users != nil {
		for _, user := range usersReq.Users {
			names = append(names, user.LoginName)
		}
	}
	if len(names) == 0 {
		return pb_user_v1.NullResponse{}, errors.New("无用户可导入")
	}
	users, _ = u.userRepo.GetUsersByLoginNames(names)
	//找出已存在的用户
	if usersReq.Users != nil {
		for _, userReq := range usersReq.Users {
			var (
				isExist bool
				id      int
			)
			for _, user := range users {
				if user.GroupID == int(userReq.GroupId) && user.LoginName == userReq.LoginName {
					isExist = true
					id = user.ID
					break
				}
			}
			if isExist {
				fmt.Println(id)
				userIsExist = append(userIsExist, models.User{
					BaseModel: models.BaseModel{
						ID: id,
					},
					UserName:  userReq.UserName,
					LoginName: userReq.LoginName,
					Password:  userReq.Password,
					GroupID:   int(userReq.GroupId),
					Mobile:    int(userReq.Mobile),
				})
				userIdsIsExist = append(userIdsIsExist, id)
			} else {
				userNotExist = append(userNotExist, models.User{
					UserName:  userReq.UserName,
					LoginName: userReq.LoginName,
					Password:  userReq.Password,
					GroupID:   int(userReq.GroupId),
					Mobile:    int(userReq.Mobile),
				})
			}
			if roleIds == nil {
				if userReq.RoleIds != nil {
					for _, index := range userReq.RoleIds {
						if index != nil {
							roleIds = append(roleIds, index.Id)
						}
					}
				}

			}
		}

	}
	//将未存在的用户插入
	tx := u.userRepo.GetTx()
	tx.Begin()
	ids, err := u.userRepo.AddUsersRepo(userNotExist, tx)

	if err != nil {
		tx.Rollback()
	}
	for _, id := range ids {
		for _, roleId := range roleIds {
			userRoles = append(userRoles, models.UserRole{
				UserID: id,
				RoleID: int(roleId),
			})
		}
	}
	err = u.userRepo.AddUserRolesRepo(userRoles, tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	userRoles = nil
	//如果覆盖写入对已存在用户修改
	tx = u.userRepo.GetTx()
	tx.Begin()
	if usersReq.IsCover == 1 {
		for _, user := range userIsExist {
			err = u.userRepo.UpdateUserByIDRepo(user, tx)
			if err != nil {
				tx.Rollback()
				break
			}
		}

		//对原有角色删除插入新的
		if err == nil {
			err = u.userRepo.DeleteUserRolesByUserId(userIdsIsExist, tx)
			if err != nil {
				tx.Rollback()
			}
			if err == nil {
				for _, id := range userIdsIsExist {
					for _, roleId := range roleIds {
						userRoles = append(userRoles, models.UserRole{
							UserID:    id,
							RoleID:    int(roleId),
						})
					}
				}
				err = u.userRepo.AddUserRolesRepo(userRoles, tx)
				if err != nil {
					tx.Rollback()
				}
			}
		}
	}
	if err == nil {
		tx.Commit()
	}

	return pb_user_v1.NullResponse{}, nil
}

func (u *userService) GetUsersSvc(ctx context.Context, data *pb_user_v1.UserQueryCondition) (*pb_user_v1.UserQueryResponse, error) {

	var userQueryResponses pb_user_v1.UserQueryResponse
	condition := &models.UserQueryByCondition{
		ID: data.Id,
		LoginName: data.LoginName,
		GroupId: data.GroupId,
		PageNum: data.PageNum,
		PageSize: data.PageSize,
	}

	querySlices,total, err := u.userRepo.GetUsersRepo(condition)
	if err != nil {
		return nil, err
	}
	userQueryResponses = pb_user_v1.UserQueryResponse{
		Total: total,
		PageNum: condition.PageNum,
		PageSize: condition.PageSize,
	}
	for _, val := range querySlices {
		_temp := &pb_user_v1.UserQueryResult{
			Id:          val.Id,
			LoginName:   val.LoginName,
			CreatedAt:   val.CreatedAt,
			UserName:    val.UserName,
			GroupName:   val.GroupName,
			GroupId:     val.GroupId,
			RoleName:    val.RoleName,
		}
		userQueryResponses.UserQueryResult = append(userQueryResponses.UserQueryResult, _temp)
	}

	return &userQueryResponses, nil
}