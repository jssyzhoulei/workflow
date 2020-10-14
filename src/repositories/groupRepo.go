package repositories

import (
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gorm.io/gorm"
)

// GroupRepoI ...
type GroupRepoI interface {
	GroupAdd(data *models.Group, tx *gorm.DB) error
	GetTx() *gorm.DB
	GroupQueryByName(name string, tx *gorm.DB) (*models.Group,error)
	QuotaAdd(data []*models.Quota, tx *gorm.DB) error
}

type groupRepo struct {
	*gorm.DB
}

func (u *groupRepo) GetTx() *gorm.DB {
	return u.Begin()
}

// NewGroupRepo ...
func NewGroupRepo(db *yorm.DB) GroupRepoI {
	return &groupRepo{
		DB: db.DB,
	}
}

// GroupAdd 添加组
func (u *groupRepo) GroupAdd(data *models.Group, tx *gorm.DB) error {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = u.DB
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
		Name: data.Name,
		ParentID: data.ParentID,
		LevelPath: parentGroup.LevelPath,
	}
	if err = db.Create(newGroupRecord).Error; err != nil {
		return err
	}

	return nil
}

// GroupQueryByName 通过组名查询组信息
func (u *groupRepo) GroupQueryByName(name string, tx *gorm.DB) (*models.Group, error) {
	var err error
	var db *gorm.DB
	if tx == nil {
		db = u.DB
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
func (u *groupRepo) QuotaAdd(data []*models.Quota, tx *gorm.DB) error {
	var err error
	err = tx.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}




