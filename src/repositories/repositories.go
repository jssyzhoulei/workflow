package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoI
	GetGroupRepo() GroupRepoInterface
	GetRoleRepo() RoleRepoI
	GetPermissionRepo() PermissionRepoInterface
}

type repo struct {
	UserRepoI
	GroupRepoInterface
	RoleRepoI
	PermissionRepoInterface
}

func (r *repo) GetPermissionRepo() PermissionRepoInterface {
	return r.PermissionRepoInterface
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoI:  NewUserRepo(db),
		GroupRepoInterface: NewGroupRepo(db),
		RoleRepoI:  NewRoleRepo(db),
		PermissionRepoInterface: NewPermissionRepo(db),
	}
}

func (r *repo) GetUserRepo() UserRepoI {
	return r.UserRepoI
}

func (r *repo) GetGroupRepo() GroupRepoInterface {
	return r.GroupRepoInterface
}
func (r *repo) GetRoleRepo() RoleRepoI {
	return r.RoleRepoI
}
