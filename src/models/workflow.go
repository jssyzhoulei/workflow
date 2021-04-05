package models

/*
工作流主表  简介  版本version会变化 始终指向最新的节点（work node）版本
*/
type WorkFLow struct {
	BaseModel
	Name        string `gorm:"column:name;type:varchar(50);comment:'工作流名称'" json:"name"`
	Description string `gorm:"column:description;type:varchar(1024);comment:'描述'" json:"description"`
	CreateID    int    `gorm:"column:create_id;type:int(10);comment:'创建者'" json:"create_id"`
	GroupID     int    `gorm:"column:group_id;type:int(10);comment:'创建者组'" json:"group_id"`
	Version     int    `gorm:"column:version;type:int(10);comment:'工作流版本'" json:"version"`
	Status      int    `gorm:"column:status;type:int(10);default:0;comment:'备用字段'" json:"status"`
}

func (WorkFLow) TableName() string {
	return "work_flow"
}

/*
 工作节点 每个工作流下都有工作节点  属于工作流子表
 本质上 并行、串行、会签、代理、条件节点都是复杂节点
 串行和会签的区别在于，串行是会签的特殊情况，必须按顺序从左到右执行，会签则无需执行顺序，两者都需要执行完所有的子节点
 复杂节点有嵌套  父节点作为容器无需指定审批人
 例外的是代理节点  代理节点有审批人  理节点的审批人可以指定代理人  代理人所在的节点即是代理节点的子节点
 // delete 错误需求 此处牵扯到一个审批逻辑，即会签节点下的子节点不能指定审批人类型为2或者3 这种情况可能导致多个会签子节点拥有统一个审批人的情况
 有发起节点，本质上是一个普通节点，结束节点不确定，有可能是动态的结束（即嵌套并行节点的每一个子节点都可能成为结束节点）
*/
type WorkNode struct {
	BaseModel
	Name        string `gorm:"column:name;type:varchar(50);comment:'工作流名称'" json:"name"`
	WorkFLowID  int    `gorm:"column:work_flow_id;type:int(10);comment:'所属工作流'" json:"work_flow_id"`
	ParentID    int    `gorm:"column:parent_id;type:int(10);comment:'父节点'" json:"parent_id"`
	SkipID      int    `gorm:"column:skip_id;type:int(10);comment:'要跳转的下一个节点'" json:"skip_id"`
	LastID      int    `gorm:"column:last_id;type:int(10);comment:'串行节点首节点外的必须字段'" json:"last_id"`
	PrincipleID int    `gorm:"column:principle_id;type:int(10);comment:'审批相关的组织/人员id'" json:"principle_id"`
	Type        int    `gorm:"column:type;type:int(2);comment:'工作流类型0 普通 1串行 2 并行 3 会签/多人拟合 4代理 5条件 6 发起节点 '" json:"type"`
	AuditType   int    `gorm:"column:audit_type;type:int(2);comment:'审批人类型0 固定人 1 组织下任何人 2 发起人上级人员 3 发起人顶级leader 4发起人下级人员 5 发起人上级组下人员 6 从工单中获取审批人 7 从共担中获取审批组'" json:"audit_type"`
	Version     int    `gorm:"column:version;type:int(10);comment:'工作流节点版本'" json:"version"`
	Status      int    `gorm:"column:status;type:int(10);default:0;comment:'备用字段'" json:"status"`
}

func (WorkNode) TableName() string {
	return "work_node"
}

type WorkNodeRequest struct {
	WorkNode
	SkipName string             `json:"skip_name"`
	Children []*WorkNodeRequest `json:"children"`
}

const (
	WorkNodeTypeGeneral = iota
	WorkNodeTypeSerial
	WorkNodeTypeParallel
	WorkNodeTypeCounterSign
	WorkNodeTypeAgent
	WorkNodeTypeCondition
	WorkNodeTypeHead

	AuditTypeAbsolute         = 0
	AuditTypeUAnyGroup        = 1
	AuditTypeLastLevel        = 2
	AuditTypeTopLeader        = 3
	AuditTypeNextLevelMembers = 4
	AuditTypeUpGroupMembers   = 5
	AuditTypeFromFormMembers  = 6
	AuditTypeFromFormGroup    = 7

	WorkInstanceFuture   = 0
	WorkInstanceComplete = 1
	WorkInstanceWaiting  = 2
	WorkInstanceSkip     = 3
	WorkInstanceReady    = 4
)

/*
 工作流实例  指向工作流节点
 状态字段  complete(完成)、waiting(复杂节点的状态)、future(未触发)、skip(并行节点被忽略的子节点)、ready(可执行的节点，可以有多个)
*/
type WorkInstance struct {
	BaseModel
	NodeID      int             `gorm:"column:node_id;type:int(10);comment:'指向的工作流节点'" json:"node_id"`
	PrincipleID int             `gorm:"column:principle_id;type:int(10);comment:'审批人员id'" json:"principle_id"`
	Status      int             `gorm:"column:status;type:int(10);default:0;comment:'状态字段'" json:"status"`
	Children    []*WorkInstance `gorm:"-" json:"children"`
}
