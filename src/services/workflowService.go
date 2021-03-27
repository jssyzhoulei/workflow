package services

import (
	"github.com/jssyzhoulei/workflow/src/models"
	"github.com/jssyzhoulei/workflow/src/repositories"
)

type WorkService struct {
	repo *repositories.WorkRepo
}

func NewWorkSvc(repo *repositories.WorkRepo) *WorkService {
	return &WorkService{repo}
}

func (ws WorkService)CreateFlow(wf *models.WorkFLow)error{
	return ws.repo.AddWorkFlow(wf)
}

func (ws WorkService)ListFlow()(interface{}, error){
	req := new(repositories.ListReq)
	return ws.repo.ListWorkFlow(req)
}

func (ws WorkService)UpdateFlow(wf *models.WorkFLow)error{
	return ws.repo.UpdateWorkFlow(wf)
}

func (ws WorkService)DelFlow(wf *models.WorkFLow)error{
	return ws.repo.DelWorkFlow(wf)
}