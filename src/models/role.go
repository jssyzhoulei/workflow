package models

type Role struct {
	BaseModel
	Name       string `gorm:"column:name;type:varchar(30);comment:'角色名'" json:"name"`
	DataPermit int    `gorm:"column:data_permit;type:int(1);comment:'数据权限1:个人；2:组织；3全部';" json:"data_permit"`
	Status     int    `gorm:"column:status;type:int(1);comment:'角色状态1:启用；2:停用';default:1" json:"status"`
}
