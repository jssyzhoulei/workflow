package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/repositories"
)

type RoleServiceI interface {
	AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error
	UpdateRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error
	DeleteRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error
}

type roleService struct {
	roleRepo repositories.RoleRepoI
}

func NewRoleService(repos repositories.RepoI) RoleServiceI {
	return &roleService{
		roleRepo: repos.GetRoleRepo(),
	}
}

func (r *roleService) AddRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error {
	err := r.roleRepo.AddRoleRepo(&role)
	if err != nil{
		return err
	}
	for _, mp := range role.MenuPerms{
		mp.RoleID = role.ID
	}
	return r.roleRepo.BatchCreateMenuPermRepo(&role.MenuPerms)
}

func (r *roleService) UpdateRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error {
	err := r.roleRepo.UpdateRoleRepo(&role.Role)
	err = r.roleRepo.DeleteMenuPermissionByRoleIDRepo(role.ID)
	if err != nil{
		return err
	}
	for _, mp := range role.MenuPerms{
		mp.ID = 0
		mp.RoleID = role.ID
	}
	return r.roleRepo.BatchCreateMenuPermRepo(&role.MenuPerms)
}

func (r *roleService) DeleteRoleSvc(ctx context.Context, role models.CreateMenuPermRequest) error{
	err := r.roleRepo.DeleteMenuPermissionByRoleIDRepo(role.ID)
	err = r.roleRepo.DeleteRoleRepo(&role.Role)
	return err
}
