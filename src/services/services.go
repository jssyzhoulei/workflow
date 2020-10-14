package services

import (
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/config"
	"gitee.com/grandeep/org-svc/utils/src/pkg/engine"
)

type ServiceI interface {
	GetUserService() UserServiceI
	GetGroupService() GroupServiceInterface
	GetRoleService() RoleServiceI
}

type service struct {
	config       *config.Config
	userService  UserServiceI
	groupService GroupServiceInterface
	roleService  RoleServiceI
}

func NewService(repo repositories.RepoI, e *engine.Engine) ServiceI {
	return &service{
		userService:  NewUserService(repo),
		groupService: NewGroupService(repo),
		roleService:  NewRoleService(repo),
		config:       e.Config,
	}
}

func (s service) GetUserService() UserServiceI {
	return s.userService
}

func (s service) GetGroupService() GroupServiceInterface {
	return s.groupService
}

func (s service) GetRoleService() RoleServiceI {
	return s.roleService
}
