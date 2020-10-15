package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoI
	GetGroupRepo() GroupRepoI
	GetRoleRepo() RoleRepoI
	GetPermissionRepo() PermissionRepoInterface
}

type repo struct {
	UserRepoI
	GroupRepoI
	RoleRepoI
	PermissionRepoInterface
}

func (r *repo) GetPermissionRepo() PermissionRepoInterface {
	return r.PermissionRepoInterface
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoI:  NewUserRepo(db),
		GroupRepoI: NewGroupRepo(db),
		RoleRepoI:  NewRoleRepo(db),
		PermissionRepoInterface: NewPermissionRepo(db),
	}
}

func (r *repo) GetUserRepo() UserRepoI {
	return r.UserRepoI
}

func (r *repo) GetGroupRepo() GroupRepoI {
	return r.GroupRepoI
}
func (r *repo) GetRoleRepo() RoleRepoI {
	return r.RoleRepoI
}
