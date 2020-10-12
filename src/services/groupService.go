package services

import (
	"gitee.com/grandeep/org-svc/src/repositories"
	"golang.org/x/net/context"
)

type GroupServiceInterface interface {

}

type GroupService struct {
	userRepo repositories.UserRepoI
}

func NewGroupService(repos repositories.RepoI) GroupServiceInterface {
	return &GroupService{
		userRepo: repos.GetUserRepo(),
	}
}

func (g GroupService) GetGroupInfo(_ context.Context, params interface{}) error {

	return  nil
}
