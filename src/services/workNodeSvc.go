package services

import (
	"github.com/jssyzhoulei/workflow/src/models"
	"gorm.io/gorm"
)

func (ws WorkService) CreateNodes(wf []*models.WorkNode) error {
	return nil
}

func (ws WorkService) parseNodes(tx *gorm.DB, parentId int, wfs []*models.WorkNode) error {
	var lastId int
	for _, i := range wfs {
		i.ParentID = parentId
		i.LastID = lastId
		err := ws.repo.AddWorkNode(tx, i)
		if err != nil {
			tx.Rollback()
			return err
		}
		lastId = i.ID
		if len(i.Children) > 0 {
			err := ws.parseNodes(tx, i.ID, i.Children)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func checkWorkNode(wfs *models.WorkNode) {
	return
}
