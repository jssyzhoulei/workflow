package services

import (
	"gitee.com/grandeep/org-svc/cmd/org-svc/engine"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/config"
)

type ServiceI interface {
	GetUserService() UserServiceInterface
	GetGroupService() GroupServiceInterface
	GetRoleService() RoleServiceI
	GetPermissionService() PermissionServiceInterface
}

type service struct {
	config       *config.Config
	userService  UserServiceInterface
	groupService GroupServiceInterface
	roleService  RoleServiceI
	permissionService PermissionServiceInterface
}

func NewService(repo repositories.RepoI, e *engine.Engine) ServiceI {
	return &service{
		userService:  NewUserService(repo, e.Config),
		groupService: NewGroupService(repo),
		roleService:  NewRoleService(repo),
		permissionService: NewPermissionService(repo),
		config:       e.Config,
	}
}

func (s service) GetUserService() UserServiceInterface {
	return s.userService
}

func (s service) GetGroupService() GroupServiceInterface {
	return s.groupService
}

func (s service) GetRoleService() RoleServiceI {
	return s.roleService
}

func (s service) GetPermissionService() PermissionServiceInterface {
	return s.permissionService
}
