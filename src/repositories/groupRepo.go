package repositories

import (
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
)

// GroupRepoInterface ...
type GroupRepoInterface interface {
	GroupAddRepo(data *models.Group, tx *gorm.DB) error
	GetTx() *gorm.DB
	GroupQueryByNameRepo(name string, tx *gorm.DB) (*models.Group, error)
	QuotaAddRepo(data []*models.Quota, tx *gorm.DB) error
	GroupQueryByConditionRepo(condition *models.GroupQueryByCondition, tx *gorm.DB) ([]*models.Group, error)
	QuotaQueryByConditionRepo(condition *models.QuotaQueryByCondition, tx *gorm.DB) ([]*models.Quota, error)
	GroupQueryWithQuotaByConditionRepo(condition *models.GroupQueryByCondition, tx *gorm.DB) ([]*models.GroupQueryWithQuotaScanRes, error)
}

type groupRepo struct {
	*gorm.DB
}

func (g *groupRepo) GetTx() *gorm.DB {
	return g.Begin()
}

// NewGroupRepo ...
func NewGroupRepo(db *yorm.DB) GroupRepoInterface {
	return &groupRepo{
		DB: db.DB,
	}
}

// GroupAdd 添加组
func (g *groupRepo) GroupAddRepo(data *models.Group, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	// 判定用户组是否存在
	var count int64
	err = db.Model(&models.Group{}).Where("name=?", data.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("创建组: %s 失败, 已经存在", data.Name)
	}

	// 查询父级数据
	var parentGroup = new(models.Group)
	err = db.Model(&models.Group{}).Where("id=?", data.ParentID).First(&parentGroup).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			parentGroup.LevelPath = "0-"
		} else {
			return err
		}
	}

	// 创建新的组
	newGroupRecord := &models.Group{
		Name:      data.Name,
		ParentID:  data.ParentID,
		LevelPath: parentGroup.LevelPath,
	}
	if err = db.Create(newGroupRecord).Error; err != nil {
		return err
	}

	return nil
}

// GroupQueryByName 通过组名查询组信息
func (g *groupRepo) GroupQueryByNameRepo(name string, tx *gorm.DB) (*models.Group, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	var record = new(models.Group)
	err = db.Model(&models.Group{}).Where("name=?", name).Find(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

// QuotaAdd 批量创建配额
func (g *groupRepo) QuotaAddRepo(data []*models.Quota, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}
	err = db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

// GroupQueryByCondition 通过条件查询组信息
func (g *groupRepo) GroupQueryByConditionRepo(condition *models.GroupQueryByCondition, tx *gorm.DB) ([]*models.Group, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	db = db.Model(&models.Group{})

	if len(condition.ID) != 0 {
		db = db.Where("id in ?", condition.ID)
	} else if len(condition.Name) != 0 {
		db = db.Where("name in ?", condition.Name)
	}
	if len(condition.ParentID) != 0 {
		db = db.Where("parent_id in ?", condition.ParentID)
	}

	var result []*models.Group
	err = db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// QuotaQueryByConditionRepo 通过条件查询配额信息
func (g *groupRepo) QuotaQueryByConditionRepo(condition *models.QuotaQueryByCondition, tx *gorm.DB) ([]*models.Quota, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	db = db.Model(&models.Quota{})

	if condition.IsShare != 0 {
		db = db.Where("is_share=?", condition.IsShare)
	}
	if condition.Type != 0 {
		db = db.Where("type=?", condition.Type)
	}
	if condition.GroupID != 0 {
		db = db.Where("group_id=?", condition.GroupID)
	}
	if condition.ResourceID != "" {
		db = db.Where("resources_id like ?", "%"+condition.ResourceID+"%")
	}

	var result = make([]*models.Quota, 0)
	err = db.Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GroupQueryWithQuotaByConditionRepo
func (g *groupRepo) GroupQueryWithQuotaByConditionRepo(condition *models.GroupQueryByCondition, tx *gorm.DB) ([]*models.GroupQueryWithQuotaScanRes, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	whereCondition := " where 1=1 "
	var conditionVal = make(map[string]interface{})

	if len(condition.ID) != 0 {
		whereCondition += " and id in @ids "
		conditionVal["ids"] = condition.ID
	} else if len(condition.Name) != 0 {
		whereCondition += " and name in @names "
		conditionVal["names"] = condition.Name
	}
	if len(condition.ParentID) != 0 {
		whereCondition += " and parent_id in @parent_ids "
		conditionVal["parent_ids"] = condition.ParentID
	}

	sqlStr := `
SELECT
	a.id,
	a.name,
	a.parent_id,
	a.level_path,
	a.created_at,
	b.is_share,
	b.resources_id,
	b.` + "`type`," + `
	b.total,
	b.used
FROM (
	SELECT
		id,
		name,
		parent_id,
		level_path,
		created_at
	FROM ` + "`group`" + fmt.Sprintf(`%s) a
	LEFT JOIN quota b ON a.id = b.group_id;
`, whereCondition)

	var result = make([]*models.GroupQueryWithQuotaScanRes, 0)
	err = db.Raw(sqlStr, conditionVal).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
