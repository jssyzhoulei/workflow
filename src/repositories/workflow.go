package repositories

import (
	"errors"
	"github.com/jssyzhoulei/workflow/src/models"
	"gorm.io/gorm"
)

type WorkRepo struct {
	*gorm.DB
}

type Page struct {
	page    int
	perPage int
}

var (
	flow = new(models.WorkFLow)
	node = new(models.WorkNode)
)

func NewRepo(db *gorm.DB) *WorkRepo {
	return &WorkRepo{db}
}

func (db *WorkRepo) AddWorkFlow(wf *models.WorkFLow) error {
	return db.Model(wf).Create(wf).Error
}

type ListReq struct {
	Name        string
	UserId      int
	GroupIdList []int
	Page
}

func (db *WorkRepo) ListWorkFlow(req *ListReq) ([]models.WorkFLow, error) {
	var res []models.WorkFLow
	if req.perPage == 0 {
		req.perPage = 10
	}
	query := db.Model(flow).Where("name like ?", "%"+req.Name+"%")
	if req.UserId != 0 {
		query = query.Where("create_id = ? ", req.UserId)
	}
	if req.GroupIdList != nil && len(req.GroupIdList) > 0 {
		query = query.Where("group_id in (?) ", req.GroupIdList)
	}
	return res, query.Limit(req.page).Offset(req.page * req.perPage).Find(&res).Error
}

func (db *WorkRepo) UpdateWorkFlow(wf *models.WorkFLow) error {
	if wf.ID == 0 {
		return errors.New("work flow record not found by update")
	}
	return db.Model(wf).Save(wf).Error
}

func (db *WorkRepo) DelWorkFlow(wf *models.WorkFLow) error {
	return db.Model(wf).Delete(wf).Error
}

func (db *WorkRepo) AddWorkNode(tx *gorm.DB, wn *models.WorkNode) error {
	if tx == nil {
		tx = db.DB
	}
	return tx.Model(wn).Create(wn).Error
}

func (db *WorkRepo) SaveWorkNode(tx *gorm.DB, wn *models.WorkNode) error {
	if tx == nil {
		tx = db.DB
	}
	return tx.Model(wn).Save(wn).Error
}
