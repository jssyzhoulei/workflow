package models

// Group 组表
type Group struct {
	BaseModel
	Name        string `gorm:"column:name;type:varchar(50);comment:'组织名称'" json:"name"`
	Description string `gorm:"column:description;type:varchar(1024);comment:'描述'" json:"description"`
	ParentID    int    `gorm:"column:parent_id;type:int(10);comment:'父级组织ID'" json:"parent_id"`
	TopID       int    `gorm:"column:top_id;type:int(10);comment:'租户ID'" json:"top_id"`
	Status      int    `gorm:"column:status;type:int(10);default:0;comment:'备用字段'" json:"status"`
}

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

// GroupQueryByCondition 组查询条件
type GroupQueryByCondition struct {
	ID       []int64  `json:"id"`
	Name     []string `json:"name"`
	ParentID []int64  `json:"parent_id"`
}

// QueryGroupsUsersResponse 查询组下的下级用户
type QueryGroupsUsersResponse struct {
	ID        int64  `json:"id"`
	UserName  string `json:"user_name"`
	LoginName string `json:"login_name"`
	GroupID   int64  `json:"group_id"`
	UserType  int    `json:"user_type"`
	Mobile    string `json:"mobile"`
}

// GroupTreeNode 组树形结构节点
type GroupTreeNode struct {
	Name     string           `json:"name"`
	ID       string           `json:"id"`
	Children []*GroupTreeNode `json:"children"`
}
