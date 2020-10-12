package repositories

import "gitee.com/grandeep/org-svc/utils/src/pkg/yorm"

type RepoI interface {
	GetUserRepo() UserRepoI
}

type repo struct {
	UserRepoI
}

func NewRepoI(db *yorm.DB) RepoI {
	return &repo{
		UserRepoI: NewUserRepo(db),
	}
}


func (r *repo) GetUserRepo() UserRepoI {
	return r.UserRepoI
}
