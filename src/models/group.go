package models

// Quota 配额表
type Quota struct {
	BaseModel
	IsShare    int          `gorm:"column:is_share;type:int(10);comment:'是否为共享 1 共享 2 独享'" json:"is_share"`
	ResourceID string       `gorm:"column:resources_id;type:varchar(255);comment:'资源组ID'" json:"resources_id"`
	Type       ResourceType `json:"type" gorm:"column:type;type:int(10);comment:'资源类型'"` // 枚举字段 ResourceType
	GroupID    int          `gorm:"column:group_id;type:int(10);comment:'组织ID'" json:"group_id"`
	Total      int          `json:"total" gorm:"column:total;type:int(10);comment:'资源总数'"`
	Used       int          `json:"used" gorm:"column:used;type:int(10);comment:'已经使用'"`
}

// TableName ...
func (q Quota) TableName() string {
	return "quota"
}

// Group 组表
type Group struct {
	BaseModel
	Name      string `gorm:"column:name;type:varchar(50);comment:'组织名称'" json:"name"`
	ParentID  int    `gorm:"column:parent_id;type:int(10);comment:'父级组织ID'" json:"parent_id"`
	LevelPath string `gorm:"column:level_path;type:varchar(255);comment:'组织等级路径'" json:"level_path"`
	Status    int    `json:"status" gorm:"column:status;type:int(10);default:0;comment:'1 已删除 0 未删除'"`
}

// TableName ...
func (g Group) TableName() string {
	return "group"
}

type ResourceType int8

func (t ResourceType) Auth() bool {
	if ResourceCpu <= t && t <= ResourceDisk {
		return true
	}
	return false
}

// 资源类型枚举
const (
	ResourceCpu ResourceType = iota + 1
	ResourceGpu
	ResourceMemory
	ResourceDisk
)

// ########################## service 参数 #################################################

// QuotaQueryByCondition 配额查询条件
type QuotaQueryByCondition struct {
	GroupID    int64  `json:"group_id"`
	Type       int    `json:"type"`
	ResourceID string `json:"resource_id"`
	IsShare    int    `json:"is_share"`
}

// GroupQueryByCondition 组查询条件
type GroupQueryByCondition struct {
	ID       []int64  `json:"id"`
	Name     []string `json:"name"`
	ParentID []int64  `json:"parent_id"`
}

// GroupQueryWithQuota 查询组和配额结果
type GroupQueryWithQuota struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ParentID    int64  `json:"parent_id"`
	TopParentID int    `json:"top_parent_id"`
	CreateTime  string `json:"create_time"`
	Quotas      []struct {
		IsShare     int    `json:"is_share"`
		ResourcesID string `json:"resources_id"`
		Type        int    `json:"type"`
		Total       int    `json:"total"`
		Used        int    `json:"used"`
	} `json:"quotas"`
}

// GroupQueryWithQuotaScanRes 查询组和配额 SQL Scan 结构体
type GroupQueryWithQuotaScanRes struct {
	ID         int64  `gorm:"column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	ParentID   int64  `gorm:"column:parent_id" json:"parent_id"`
	LevelPath  string `gorm:"column:level_path" json:"level_path"`
	CreateTime string `gorm:"column:id" json:"create_time"`
	IsShare    int    `gorm:"column:is_share" json:"is_share"`
	ResourceID string `gorm:"column:resources_id" json:"resources_id"`
	Type       int    `gorm:"column:type" json:"type"`
	Total      int    `gorm:"column:total" json:"total"`
	Used       int    `gorm:"column:used" json:"used"`
}

// GroupUpdateRequest 组信息更新请求
type GroupUpdateRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID *int64 `json:"parent_id"`
}

// QuotaUpdateRequest 配额更新请求
type QuotaUpdateRequest struct {
	GroupID     int64  `json:"group_id"`
	IsShare     int64  `json:"is_share"`
	ResourcesID string `json:"resources_id"`
	QuotaType   int64  `json:"quota_type"`
	Total       int64  `json:"total"`
	Used        int64  `json:"used"`
}
