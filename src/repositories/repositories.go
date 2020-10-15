package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoInterface
	GetGroupRepo() GroupRepoInterface
	GetRoleRepo() RoleRepoI
	GetPermissionRepo() PermissionRepoInterface
}

type repo struct {
	UserRepoInterface
	GroupRepoInterface
	RoleRepoI
	PermissionRepoInterface
}

func (r *repo) GetPermissionRepo() PermissionRepoInterface {
	return r.PermissionRepoInterface
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoInterface:  NewUserRepo(db),
		GroupRepoInterface: NewGroupRepo(db),
		RoleRepoI:  NewRoleRepo(db),
		PermissionRepoInterface: NewPermissionRepo(db),
	}
}

func (r *repo) GetUserRepo() UserRepoInterface {
	return r.UserRepoInterface
}

func (r *repo) GetGroupRepo() GroupRepoInterface {
	return r.GroupRepoInterface
}
func (r *repo) GetRoleRepo() RoleRepoI {
	return r.RoleRepoI
}
