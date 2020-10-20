package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
)

type RoleServiceI interface {
	AddRoleSvc(ctx context.Context, role *models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error)
	UpdateRoleSvc(ctx context.Context, role *models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error)
	DeleteRoleSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error)
	QueryRoleSvc(ctx context.Context, id int) (*models.CreateMenuPermRequest, error)
	QueryRolesSvc(ctx context.Context, page *pb_user_v1.RolePageRequestProto) (*pb_user_v1.RolePageRequestProto, error)
}

type roleService struct {
	roleRepo repositories.RoleRepoI
}

func NewRoleService(repos repositories.RepoI) RoleServiceI {
	return &roleService{
		roleRepo: repos.GetRoleRepo(),
	}
}

func (r *roleService) AddRoleSvc(ctx context.Context, role *models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	err := r.roleRepo.AddRoleRepo(role)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	for _, mp := range role.MenuPerms {
		mp.RoleID = role.ID
	}
	return pb_user_v1.NullResponse{}, r.roleRepo.BatchCreateMenuPermRepo(&role.MenuPerms)
}

func (r *roleService) UpdateRoleSvc(ctx context.Context, role *models.CreateMenuPermRequest) (pb_user_v1.NullResponse, error) {
	err := r.roleRepo.UpdateRoleRepo(&role.Role)
	err = r.roleRepo.DeleteMenuPermissionByRoleIDRepo(role.ID)
	if err != nil {
		return pb_user_v1.NullResponse{}, err
	}
	for _, mp := range role.MenuPerms {
		mp.ID = 0
		mp.RoleID = role.ID
	}
	return pb_user_v1.NullResponse{}, r.roleRepo.BatchCreateMenuPermRepo(&role.MenuPerms)
}

func (r *roleService) DeleteRoleSvc(ctx context.Context, id int) (pb_user_v1.NullResponse, error) {
	err := r.roleRepo.DeleteMenuPermissionByRoleIDRepo(id)
	role := models.Role{}
	role.ID = id
	err = r.roleRepo.DeleteRoleRepo(&role)
	return pb_user_v1.NullResponse{}, err
}

func (r *roleService) QueryRoleSvc(ctx context.Context, id int) (*models.CreateMenuPermRequest, error) {
	return r.roleRepo.RoleDetailRepo(id, 0)
}

func (r *roleService) QueryRolesSvc(ctx context.Context, page *pb_user_v1.RolePageRequestProto) (*pb_user_v1.RolePageRequestProto, error) {
	return r.roleRepo.ListRolesRepo(page, 0)
}
