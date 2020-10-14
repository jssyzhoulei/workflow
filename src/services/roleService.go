package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/repositories"
)

type RoleServiceI interface {
	AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) (models.Role, error)
}

type roleService struct {
	roleRepo repositories.RoleRepoI
}

func NewRoleService(repos repositories.RepoI) RoleServiceI {
	return &roleService{
		roleRepo: repos.GetRoleRepo(),
	}
}

func (r *roleService) AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) (models.Role, error) {
	err := r.roleRepo.AddRoleRepo(&role)
	if err != nil{
		return role.Role, err
	}
	for _, mp := range role.MenuPerms{
		mp.RoleID = role.ID
	}
	return role.Role, r.roleRepo.BatchCreateMenuPermRepo(&role.MenuPerms)
}
