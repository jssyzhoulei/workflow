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
	QuotaUpdateRepo(data []*models.QuotaUpdateRequest, groupID int64, tx *gorm.DB) error
	GroupListWithChangedLevelPathRepo(groupID int64, tx *gorm.DB) ([]*models.Group, error)
	GroupDeleteRepo(id int64, tx *gorm.DB) error
	QueryGroupIDAndSubGroupsID(groupID int64, tx *gorm.DB) ([]int64, error)
	SetGroupQuotaUsedRepo(data *models.SetGroupQuotaRequest, tx *gorm.DB) error
	GetAllGroup() []models.Group
	UpdateQuotaResourceID(groupID int64, resourceIDMap map[int64]string, tx *gorm.DB) error
	QueryQuota(groupID int64, tx *gorm.DB) (*models.QueryQuota, error)
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

// GroupAddRepo 添加组
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
		if err.Error() == "record not found" {
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
	err = db.Find(&result).Error
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
	a.description,
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
		description,
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

	// 获取组原本的信息
	oldGroup, err := g.GroupQueryByIDRepo(data.ID, nil)
	if err != nil {
		return err
	}

	if oldGroup.ID == 0 {
		return errors.New("组信息被标记为删除或组不存在")
	}

	if data.ParentID != nil && int64(oldGroup.ParentID) != *data.ParentID {

		// 更新父级ID时,不允许跨越顶级组ID更新
		if oldGroup.ParentID == 0 {
			return errors.New("顶级组不允许执行变更父级操作")
		} else {
			// 不允许更新含有下级组的父级
			res, err := g.QueryGroupIDAndSubGroupsID(data.ID, db)
			if err != nil {
				return err
			}
			if len(res) > 1 {
				return errors.New("包含子级不允许更新父级信息")
			}

			// 没有子级parent为0, 直接提高层级至顶级组
			if *data.ParentID == 0 {
				updateColumnMap["level_path"] = "0-"
			} else {
				// 获取新的父级组ID
				newParentGroup, err := g.GroupQueryByIDRepo(*data.ParentID, db)
				if err != nil {
					return err
				}

				// 获取顶级组ID
				oriTopGroupID := strings.Split(oldGroup.LevelPath, "-")[1]

				// 获取新的父级组的顶级组
				var newTopGroupID string
				// 如果新的父级就是顶级组那么直接对比该组的ID, 否则获取新的父级组的顶级组ID
				if newParentGroup.ParentID == 0 {
					// 这里LevelPath应该是这个样子: 0-
					// 所以直接设置为当前新父级的ID
					newTopGroupID = strconv.Itoa(newParentGroup.ID)
				} else {
					// 这里LevelPath应该是这个样子: 0-ID-ID
					newTopGroupID = strings.Split(newParentGroup.LevelPath, "-")[1]
				}

				// 判断是否跨越顶级组
				if oriTopGroupID != newTopGroupID {
					return errors.New("不允许跨越顶级组更新其父级ID")
				}

				updateColumnMap["level_path"] = newParentGroup.LevelPath + strconv.FormatInt(*data.ParentID, 10) + "-"
			}

			updateColumnMap["parent_id"] = data.ParentID
		}

	}

	if data.Description != "" {
		updateColumnMap["description"] = data.Description
	}

	err = db.Model(&models.Group{}).Where("id=?", data.ID).Updates(updateColumnMap).Error
	if err != nil {
		return err
	}

	return nil
}

// QuotaUpdateRepo 配额信息更新
func (g *groupRepo) QuotaUpdateRepo(_data []*models.QuotaUpdateRequest,  groupID int64, tx *gorm.DB) error {
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	// 增加组信息校验,被删除的组,无法修改其配额数据
	group, err := g.GroupQueryByIDRepo(groupID, nil)
	if err != nil {
		return errors.New("查询组信息错误: " + err.Error())
	}

	if group.Status == 1 {
		return errors.New("组已删除,无法修改数据")
	}

	l := len(_data)
	for i := 0; i < l; i++ {
		data := _data[i]

		if data.GroupID == 0 || (data.IsShare == 0 && data.QuotaType != int64(models.ResourceDisk)) {
			return errors.New("检测到空值: group_id, is_share 全部为必传参数")
		}

		if !models.ResourceType.Auth(models.ResourceType(data.QuotaType)) {
			return errors.New("资源类型不存在")
		}

		// 插入或者更新
		var num int64
		err = db.Model(&models.Quota{}).Where("group_id=? and is_share=? and type=?", data.GroupID,
			data.IsShare, data.QuotaType).Count(&num).Error
		if err != nil {
			return err
		}
		if num == 0 {
			// 当配额不存在时,创建对应的配额
			var  data = &models.Quota{
				IsShare:    int(data.IsShare),
				ResourceID: data.ResourcesID,
				Type:       models.ResourceType(data.QuotaType),
				GroupID:    int(data.GroupID),
				Total:      int(data.Total),
				Used:       0,
			}
			err = db.Create(&data).Error
			if err != nil {
				return errors.New("QuotaUpdateRepo 创建配额错误: " + err.Error())
			}
		} else {
			// 配额存在时更新总量
			updateColumnMap := map[string]interface{} {
				"total": data.Total,
			}

			err = db.Model(&models.Quota{}).Where("group_id=? and is_share=? and type=?", data.GroupID,
				data.IsShare, data.QuotaType).Updates(updateColumnMap).Error
			if err != nil {
				return err
			}
		}
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

	var group = new(models.Group)
	err = db.Model(&models.Group{}).Where("id=?", id).First(&group).Error
	if err != nil {
		return err
	}

	updateColumnMap := map[string]interface{} {
		"name": group.Name + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "_deleted",
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
	var levelPath string
	if groupID == 0 {
		levelPath = "%" + strconv.FormatInt(groupID, 10) + "-%"
	} else {
		levelPath = "%-" + strconv.FormatInt(groupID, 10) + "-%"
	}

	err = db.Model(&models.Group{}).Select("id").Where("level_path like ? or id=? and status=0", levelPath, groupID).Find(&groupIDs).Error
	if err != nil {
		return nil, err
	}

	if len(groupIDs) < 1 {
		return nil, errors.New("组查询记录为空")
	}

	return groupIDs, nil
}

// SetGroupQuotaUsed 设置组已使用配额
func (g *groupRepo) SetGroupQuotaUsedRepo(data *models.SetGroupQuotaRequest, tx *gorm.DB) error {
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	if data.GroupID == 0 || (data.IsShare == 0 && data.QuotaType != int64(models.ResourceDisk)) {
		return errors.New("检测到空值: group_id, is_share 全部为必传参数")
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

	quotaTableName := models.Quota{}.TableName()

	sqlStr := "update %s set used=used+%d where group_id=%d and is_share=%d and type=%d"

	fullSql := fmt.Sprintf(sqlStr, quotaTableName, data.Used, data.GroupID, data.IsShare, data.QuotaType)

	err = db.Exec(fullSql).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllGroup 获取所有的组
func (g *groupRepo) GetAllGroup() []models.Group {
	var (
		groups []models.Group
	)
	g.Find(&groups)
	return groups
}

// UpdateQuotaResourceID 更新配额资源ID
// @param resourceIDMap map[int64]string key为share 值为资源组ID字符串
// 流程: 获取用户信息 -> 判断是否需要更新资源组 -> 判断是否存在用户 -> 查询组下资源组个数 -> 判断是否需要删除 -> 判断是否存在用户
// 任一判断失败,则返回err
func (g *groupRepo) UpdateQuotaResourceID(groupID int64, resourceIDMap map[int64]string, tx *gorm.DB) error {
	var db *gorm.DB
	var err error
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}

	// 校验组下是否存在用户
	var userCount int64
	err = db.Model(&models.User{}).Where("group_id=?", groupID).Count(&userCount).Error
	if err != nil {
		return err
	}

	subGroupIDs, err := g.QueryGroupIDAndSubGroupsID(groupID, tx)
	if err != nil {
		return err
	}
	var subGroupIDArray = make([]int64, 0, len(subGroupIDs) - 1)
	for  i:=0;i<len(subGroupIDs);i++ {
		if subGroupIDs[i] == groupID {
			continue
		}
		subGroupIDArray = append(subGroupIDArray, subGroupIDs[i])
	}

	for share, resourceIDStr := range resourceIDMap {

		// 查询原资源组ID
		var oriResourcesIDStr string
		err = db.Model(&models.Quota{}).Select("resources_id").Where("group_id=? and is_share=? and type<>4", groupID,
			share).First(&oriResourcesIDStr).Error
		if err != nil {
			return err
		}

		// 组ID + 共享类型 查询资源组字段,相同则不操作
		if oriResourcesIDStr == resourceIDStr {
			continue
		}

		if userCount != 0 || len(subGroupIDArray) != 0 {
			return errors.New("组下包含用户或子组,无法更新资源组")
		}

		// 更新资源组信息
		updateColumnMap := map[string]interface{}{
			"resources_id": resourceIDStr,
		}

		err = db.Model(&models.Quota{}).Where("group_id=? and is_share=? and type<>4", groupID,
			share).Updates(updateColumnMap).Error
		if err != nil {
			return err
		}
	}

	// 查询当前组包含有几个资源组
	var idSlice = make([]string, 0, 2)
	err = db.Model(&models.Quota{}).Select("resources_id").Where("group_id=? and type<>4", groupID).Distinct("resources_id").Find(&idSlice).Error
	if err != nil {
		return errors.New("查询资源ID时错误: " + err.Error())
	}

	// 如果已经存在的资源组数量 小于 传入的
	// 例如: 已经存在 2 传入 1 表示配额表含有待删除的配额数据
	if len(idSlice) > len(resourceIDMap) {
		if userCount != 0 || len(subGroupIDArray) != 0{
			return errors.New("组下包含用户或子组,无法删除资源组")
		}
		// 进行删除
		var resIDs = make([]string, 0, 1)
		for _, id := range resourceIDMap {
			resIDs = append(resIDs, id)
		}

		err = db.Unscoped().Where("group_id=? and resources_id not in ? and type<>4", groupID, resIDs).Delete(&models.Quota{}).Error
		if err != nil {
			return errors.New("删除资源组时错误: " + err.Error())
		}
	}

	return nil
}

// QueryQuota 查询配额
func (g *groupRepo) QueryQuota(groupID int64, tx *gorm.DB) (*models.QueryQuota, error) {
	var db *gorm.DB
	if tx == nil {
		db = g.DB
	} else {
		db = tx
	}
	var res = make([]*models.Quota, 0, 4)
	err := db.Model(&models.Quota{}).Where("group_id=?", groupID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	var result = new(models.QueryQuota)
	var cache = make(map[int]*models.QueryQuotaInfo)
	l := len(res)
	for i := 0; i < l; i++ {
		data := res[i]
		if _, ok := cache[data.IsShare]; !ok {
			cache[data.IsShare] = new(models.QueryQuotaInfo)
		}

		switch data.Type {
		case models.ResourceCpu:
			cache[data.IsShare].ResourcesGroupId = data.ResourceID
			cache[data.IsShare].CpuTotal = data.Total
			cache[data.IsShare].CpuUsed = data.Used
		case models.ResourceGpu:
			cache[data.IsShare].GpuTotal = data.Total
			cache[data.IsShare].GpuUsed = data.Used
		case models.ResourceMemory:
			cache[data.IsShare].MemoryTotal = data.Total
			cache[data.IsShare].MemoryUsed = data.Used
		case models.ResourceDisk:
			result.DiskQuotaTotal = data.Total
			result.DiskQuotaUsed = data.Used
		}
	}
	if _, ok := cache[1]; ok {
		result.ShareQuota = cache[1]
	} else {
		result.ShareQuota = new(models.QueryQuotaInfo)
	}
	if _, ok := cache[2]; ok {
		result.NonShareQuota = cache[2]
	} else {
		result.NonShareQuota = new(models.QueryQuotaInfo)
	}

	return result, nil
}