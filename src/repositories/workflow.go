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

func (wr *WorkRepo) AddWorkFlow(wf *models.WorkFLow) error {
	return wr.Model(wf).Create(wf).Error
}

type ListReq struct {
	Name        string
	UserId      int
	GroupIdList []int
	Page
}

func (wr *WorkRepo) ListWorkFlow(req *ListReq) ([]models.WorkFLow, error) {
	var res []models.WorkFLow
	if req.perPage == 0 {
		req.perPage = 10
	}
	query := wr.Model(flow).Where("name like ?", "%"+req.Name+"%")
	if req.UserId != 0 {
		query = query.Where("create_id = ? ", req.UserId)
	}
	if req.GroupIdList != nil && len(req.GroupIdList) > 0 {
		query = query.Where("group_id in (?) ", req.GroupIdList)
	}
	return res, query.Limit(req.page).Offset(req.page * req.perPage).Find(&res).Error
}

func (wr *WorkRepo) UpdateWorkFlow(wf *models.WorkFLow) error {
	if wf.ID == 0 {
		return errors.New("work flow record not found by update")
	}
	return wr.Model(wf).Save(wf).Error
}

func (wr *WorkRepo) DelWorkFlow(wf *models.WorkFLow) error {
	// 关联删除work node 也可以不删
	return wr.Model(wf).Delete(wf).Error
}

func (wr *WorkRepo) AddWorkNode(tx *gorm.DB, wn *models.WorkNode) error {
	if tx == nil {
		tx = wr.DB
	}
	return tx.Model(wn).Create(wn).Error
}

func (wr *WorkRepo) SaveOrCreateWorkNode(tx *gorm.DB, wn *models.WorkNode) error {

	if wn.ID == 0 {
		return wr.AddWorkNode(tx, wn)
	}
	return wr.UpdateWorkNode(tx, wn)
}

// 局部更新
func (wr *WorkRepo) UpdateWorkNode(tx *gorm.DB, wn *models.WorkNode) error {
	if tx == nil {
		tx = wr.DB
	}
	return tx.Model(wn).Updates(wn).Error
}

// 全量更新
func (wr *WorkRepo) SaveWorkNode(tx *gorm.DB, wn *models.WorkNode) error {
	if tx == nil {
		tx = wr.DB
	}
	return tx.Model(wn).Save(wn).Error
}

func (wr *WorkRepo) ListWorkNode(flowId int) ([]models.WorkNode, error) {
	var res []models.WorkNode

	return res, wr.Model(node).Where("work_flow_id = ?", flowId).Find(&res).Error
}

func (wr *WorkRepo) DelWorkNode(ids []int) error {
	var res models.WorkNode

	return wr.Model(node).Where("id in (?)", ids).Delete(&res).Error
}
