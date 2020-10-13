package models

// 配额表
type Quota struct {
	BaseModel
	IsShare    int          `gorm:"column:is_share;type:int(10);comment:'是否为共享 1 共享 2 独享'" json:"is_share"`
	ResourceID int          `gorm:"column:resources_id;type:int(10);comment:'资源组ID'" json:"resources_id"`
	Type       ResourceType `json:"type" gorm:"column:type;type:int(10);comment:'资源类型'"`
	GroupID    int          `gorm:"column:group_id;type:int(10);comment:'组织ID'" json:"group_id"`
	Total      int          `json:"total" gorm:"column:total;type:int(10);comment:'资源总数'"`
	Used       int          `json:"used" gorm:"column:used;type:int(10);comment:'已经使用'"`
}

func (q Quota) TableName() string {
	return "quota"
}

type Group struct {
	BaseModel
	Name      string `gorm:"column:name;type:varchar(50);comment:'组织名称'" json:"name"`
	ParentID  int    `gorm:"column:parent_id;type:int(10);comment:'父级组织ID'" json:"parent_id"`
	LevelPath string `gorm:"column:level_path;type:varchar(255);comment:'组织等级路径'" json:"level_path"`
}

func (g Group) TableName() string {
	return "group"
}

//资源类型
type ResourceType int8

const (
	RESOURCE_CPU ResourceType = iota + 1
	RESOURCE_GPU
	RESOURCE_MEMORY
	RESOURCE_DISK
)
