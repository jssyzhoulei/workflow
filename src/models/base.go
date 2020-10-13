package models

import "time"

type BaseModel struct {
	ID            int        `gorm:"column:id;primary_key;not null;auto_increment" json:"id"`
	CreateUserID  int        `gorm:"column:create_user_id;type:int(10);comment:'创建用户id'" json:"create_user_id"`
	CreateGroupID int        `gorm:"column:create_group_id;type:int(10);comment:'创建用户组'" json:"create_group_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Remark        string     `gorm:"column:remark;type:text;comment:'备注'" json:"remark"`
	Extend        string     `gorm:"column:extend;type:text;comment:'扩展字段'" json:"extend"`
}
