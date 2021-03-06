package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID            int            `gorm:"column:id;primary_key;not null;auto_increment" json:"id"`
	CreateUserID  int            `gorm:"column:create_user_id;type:int(10);comment:'创建用户id'" json:"-"`
	CreateGroupID int            `gorm:"column:create_group_id;type:int(10);comment:'创建用户组'" json:"-"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	Remark        string         `gorm:"column:remark;type:text;comment:'备注'" json:"remark"`
	Extend        string         `gorm:"column:extend;type:text;comment:'扩展字段'" json:"extend"`
}

type Page struct {
	Total     int64       `json:"total"`
	TotalPage float64     `json:"total_page"`
	PageSize  int         `json:"page_size"`
	PageNum   int         `json:"page_num"`
	Data      interface{} `json:"data"`
}

// mysql -uroot -hlocalhost -p123456
// CREATE DATABASE `workflow` CHARACTER SET utf8mb4；
//sudo docker run -d -p 3306:3306 -v /home/yangyin/docker/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD="123456" mysql:5.7
//sudo docker run -d -p 3306:3306 -v /Users/sm2072/docker/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD="123456" mysql:5.7