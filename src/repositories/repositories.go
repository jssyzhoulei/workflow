package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoI
	GetRoleRepo() RoleRepoI
}

type repo struct {
	UserRepoI
	RoleRepoI
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoI: NewUserRepo(db),
		RoleRepoI: NewRoleRepo(db),
	}
}

func (r *repo) GetUserRepo() UserRepoI {
	return r.UserRepoI
}

func (r *repo) GetRoleRepo() RoleRepoI {
	return r.RoleRepoI
}
