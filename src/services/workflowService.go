package services

import (
	"github.com/jssyzhoulei/workflow/src/models"
	"github.com/jssyzhoulei/workflow/src/repositories"
	"gorm.io/gorm"
)

type WorkService struct {
	repo *repositories.WorkRepo
}

func NewWorkSvc(repo *repositories.WorkRepo) *WorkService {
	return &WorkService{repo}
}

func (ws WorkService) CreateFlow(wf *models.WorkFLow) error {
	return ws.repo.AddWorkFlow(wf)
}

func (ws WorkService) ListFlow() (interface{}, error) {
	req := new(repositories.ListReq)
	return ws.repo.ListWorkFlow(req)
}

func (ws WorkService) UpdateFlow(wf *models.WorkFLow) error {
	return ws.repo.UpdateWorkFlow(wf)
}

func (ws WorkService) DelFlow(wf *models.WorkFLow) error {
	return ws.repo.DelWorkFlow(wf)
}

func (ws WorkService) CreateNodes(wf []*models.WorkNodeRequest) error {
	ns := NewNodeSvc(ws.repo)
	// 此方法中事物commit、rollback自动进行
	return ws.repo.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		return ns.parseNodes(tx, 0, wf)
	})
}
