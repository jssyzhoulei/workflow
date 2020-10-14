package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoInterface
	GetGroupRepo() GroupRepoI
	GetRoleRepo() RoleRepoI
}

type repo struct {
	UserRepoInterface
	GroupRepoI
	RoleRepoI
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoInterface:  NewUserRepo(db),
		GroupRepoI: NewGroupRepo(db),
		RoleRepoI:  NewRoleRepo(db),
	}
}

func (r *repo) GetUserRepo() UserRepoInterface {
	return r.UserRepoInterface
}

func (r *repo) GetGroupRepo() GroupRepoI {
	return r.GroupRepoI
}
func (r *repo) GetRoleRepo() RoleRepoI {
	return r.RoleRepoI
}
