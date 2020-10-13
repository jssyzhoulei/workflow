package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoI
	GetGroupRepo() GroupRepoI
	GetRoleRepo() RoleRepoI
}

type repo struct {
	UserRepoI
	GroupRepoI
	RoleRepoI
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoI:  NewUserRepo(db),
		GroupRepoI: NewGroupRepo(db),
		RoleRepoI:  NewRoleRepo(db),
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
