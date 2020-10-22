package repositories

import (
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// GroupRepoInterface ...
type GroupRepoInterface interface {
	GroupAddRepo(data *models.Group, tx *gorm.DB) error
	GetTx() *gorm.DB
	GroupQueryByNameRepo(name string, tx *gorm.DB) (*models.Group, error)
	GroupQueryByIDRepo(id int64, tx *gorm.DB) (*models.Group, error)
	QuotaAddRepo(data []*models.Quota, tx *gorm.DB) error
	GroupQueryByConditionRepo(condition *models.GroupQueryByCondition, tx *gorm.DB) ([]*models.Group, error)
	QuotaQueryByConditionRepo(condition *models.QuotaQueryByCondition, tx *gorm.DB) ([]*models.Quota, error)
	GroupQueryWithQuotaByConditionRepo(condition *models.GroupQueryByCondition, tx *gorm.DB) ([]*models.GroupQueryWithQuotaScanRes, error)
	GroupUpdateRepo(data *models.GroupUpdateRequest, tx *gorm.DB) error
	QuotaUpdateRepo(data *models.QuotaUpdateRequest, tx *gorm.DB) error
	GroupListWithChangedLevelPathRepo(groupID int64, tx *gorm.DB) ([]*models.Group, error)
	GroupDeleteRepo(id int64, tx *gorm.DB) error
	QueryGroupIDAndSubGroupsID(groupID int64, tx *gorm.DB) ([]int64, error)

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

	// 判定组是否存在
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
	var parentIsNotExist = false
	var levelPath string
	err = db.Model(&models.Group{}).Where("id=?", data.ParentID).First(&parentGroup).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			parentIsNotExist = true
			levelPath = "0-"
		} else {
			return err
		}
	}

	if !parentIsNotExist {
		levelPath = parentGroup.LevelPath + strconv.FormatInt(int64(parentGroup.ID), 10) + "-"
	}

	// 创建新的组
	newGroupRecord := &models.Group{
		Name:      data.Name,
		ParentID:  data.ParentID,
		LevelPath: levelPath,
		NameSpace: data.NameSpace,
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

// GroupQueryByIDRepo 通过组ID查询组信息
func (g *groupRepo) GroupQueryByIDRepo(id int64, tx *gorm.DB) (*models.Group, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	var record = new(models.Group)
	err = db.Model(&models.Group{}).Where("id=?", id).Find(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

// QuotaAddRepo 批量创建配额
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

// GroupQueryByConditionRepo 通过条件查询组信息
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

// GroupQueryWithQuotaByConditionRepo 通过条件查询组及其配额信息接口
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
	// 过滤软删除的数据
	whereCondition += " and status=0 "

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

// GroupUpdateRepo 更新组信息
func (g *groupRepo) GroupUpdateRepo(data *models.GroupUpdateRequest, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	if data.ID == 0 {
		return errors.New("组ID必须传递")
	}

	updateColumnMap := make(map[string]interface{})
	if data.Name != "" {
		updateColumnMap["name"] = data.Name
	}

	if data.ParentID != nil {
		updateColumnMap["parent_id"] = data.ParentID

		oldGroup, err := g.GroupQueryByIDRepo(data.ID, nil)
		if err != nil {
			return err
		}
		oldLevelPath := oldGroup.LevelPath
		res := strings.Split(oldLevelPath, "-")
		res[len(res) - 2] = strconv.FormatInt(*data.ParentID, 10)
		newLevelPath := strings.Join(res, "-")
		updateColumnMap["level_path"] = newLevelPath
	}

	err = db.Model(&models.Group{}).Where("id=?", data.ID).Updates(updateColumnMap).Error
	if err != nil {
		return err
	}

	return nil
}

// QuotaUpdateRepo 配额信息更新
func (g *groupRepo) QuotaUpdateRepo(data *models.QuotaUpdateRequest, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	if data.GroupID == 0 || data.IsShare == 0 || data.ResourcesID == "" {
		return errors.New("检测到空值: group_id, is_share, resources_id 全部为必传参数")
	}

	if !models.ResourceType.Auth(models.ResourceType(data.QuotaType)) {
		return errors.New("资源类型不存在")
	}

	// 增加组信息校验,被删除的组,无法修改其配额数据
	group, err := g.GroupQueryByIDRepo(data.GroupID, nil)
	if err != nil {
		return errors.New("查询组信息错误: " + err.Error())
	}
	
	if group.Status == 1 {
		return errors.New("组已删除,无法修改数据")
	}

	updateColumnMap := map[string]interface{} {
		"total": data.Total,
		"used": data.Used,
	}

	err = db.Model(&models.Quota{}).Where("group_id=? and is_share=? and resources_id=? and type=?", data.GroupID,
		data.IsShare, data.ResourcesID, data.QuotaType).Updates(updateColumnMap).Error
	if err != nil {
		return err
	}

	return nil
}

// GroupDeleteRepo 组删除
func (g *groupRepo) GroupDeleteRepo(id int64, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	updateColumnMap := map[string]interface{} {
		"status": 1,
		"deleted_at": time.Now().Format("2006-01-02 15:04:05"),
	}


	err = db.Model(&models.Group{}).Where("id=?", id).Updates(updateColumnMap).Error
	if err != nil {
		return err
	}
	return nil
}

// GroupListWithChangedLevelPathRepo 通过ID查询组及其下级组信息,查询结果的 levelPath 会从顶级组开始
// 此方法被 services.GroupTreeQuerySvc 使用,用于生成 group 树形数据
func (g *groupRepo) GroupListWithChangedLevelPathRepo(groupID int64, tx *gorm.DB) ([]*models.Group, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	// 查询组ID为0时,其level_path不需要处理,取消子查询
	var subQueryLevelPath = "level_path"
	if groupID != 0 {
		subQueryLevelPath = fmt.Sprintf("substring(level_path,(select LENGTH(level_path) from `group` where id=%d and status = 0) + 1) as level_path", groupID)
	}

	sqlStr := fmt.Sprintf(`
select id,name,parent_id,
	%s
from `, subQueryLevelPath) + "`group`" + ` where level_path like ? and status = 0 order by id;
`
	var lpStr = "%"+strconv.Itoa(int(groupID)) + "-%"
	var result = make([]*models.Group, 0)
	err = db.Raw(sqlStr, lpStr).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryGroupIDAndSubGroupsID 查询组下包含其子组的ID
func (g *groupRepo) QueryGroupIDAndSubGroupsID(groupID int64, tx *gorm.DB) ([]int64, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	var groupIDs []int64
	levelPath := "%" + strconv.FormatInt(groupID, 10) + "-%"
	err = db.Model(&models.Group{}).Select("id").Where("level_path like ? or id=? and status=0", levelPath, groupID).Find(&groupIDs).Error
	if err != nil {
		return nil, err
	}

	if len(groupIDs) < 1 {
		return nil, errors.New("组查询记录为空")
	}

	return groupIDs, nil
}