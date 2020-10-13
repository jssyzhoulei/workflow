package models

import "time"

type User2 struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

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

const (
	// 算力云
	MODULE_TOP = iota
	// 基础模块
	MODULE_BASIC
	// 标注模块
	MODULE_ANNOTATION
	MODULE_TRAINING
	MODULE_DEVELOP
	MODULE_SERVICE
)

type Role struct {
	BaseModel
	Name       string `gorm:"column:name;type:varchar(30);comment:'角色名'" json:"name"`
	GroupName  string `gorm:"column:group_name;type:varchar(30);comment:'组织'" json:"group_name"`
	Module     int    `gorm:"column:module;type:int(10);comment:'所属模块'" json:"module"`
	GroupID    int    `gorm:"column:group_id;type:int(1);comment:'组织id';" json:"group_id"`
	DataPermit int    `gorm:"column:data_permit;type:int(1);comment:'数据权限0:个人；1:组织；2全部';" json:"data_permit"`
	Status     int    `gorm:"column:status;type:int(1);comment:'角色状态0:启用；1:停用';" json:"status"`
}

type MenuComponents struct {
	BaseModel
	Name     string `gorm:"column:name;type:varchar(128);not null;unique_index;comment:'菜单名称'" json:"name"`
	ParentID int    `gorm:"column:parent_id;type:int(10);comment:'父组件'" json:"parent_id"`
	Module   int    `gorm:"column:module;type:int(10);comment:'所属模块'" json:"module"`
	Order    int    `gorm:"column:order;type:int(10);comment:'组件次序'" json:"order"`
	Version  int    `gorm:"column:version;type:int(10);comment:'版本号'" json:"version"`
	Path     string `gorm:"column:path;type:varchar(512);not null;comment:'对应url path'" json:"path"`
	Status   int    `gorm:"column:status;type:int(2);not null;comment:'菜单状态 0 启用 1 未启用'" json:"status"`
}

type MenuPermission struct {
	BaseModel
	Version     int `gorm:"column:version;type:int(10);comment:'版本号'" json:"version"`
	RoleID      int `gorm:"column:role_id;type:int(10);comment:'角色id'" json:"role_id"`
	ComponentID int `gorm:"column:component_id;type:int(10);comment:'组件id'" json:"component_id"`
	Order       int `gorm:"column:order;type:int(10);comment:'属性次序'" json:"order"`
	Create      int `gorm:"column:create;type:int(10);comment:'创建 0 未选中 1 选中'" json:"create"`
	Edit        int `gorm:"column:edit;type:int(10);comment:'编辑 0 未选中 1 选中'" json:"edit"`
	Delete      int `gorm:"column:delete;type:int(10);comment:'删除 0 未选中 1 选中'" json:"delete"`
	Status      int `gorm:"column:status;type:int(2);not null;comment:'属性状态：0启用 1未启用" json:"status"`
}

type User struct {
	BaseModel
	UserName  string `gorm:"column:user_name;type:varchar(50);comment:'用户名'" json:"user_name"`
	LoginName string `gorm:"column:login_name;type:varchar(50);comment:'登录名'" json:"login_name"`
	Password  string `gorm:"column:password;type:varchar(50);comment:'密码'" json:"password"`
	GroupID   int    `gorm:"column:group_id;type:int(10);comment:'所属组织'" json:"group_id"`
	Mobile    int    `gorm:"column:mobile;type:int(10);comment:'手机号'" json:"mobile"`
	UserType  int    `gorm:"column:user_type;type:int(10);comment:'用户类型0 普通用户 1 管理员 2超级管理员'" json:"user_type"`
}

type UserRole struct {
	BaseModel
	UserID   int    `gorm:"column:user_id;type:int(10);comment:'所属组织'" json:"user_id"`
	RoleID   int    `gorm:"column:role_id;type:int(10);comment:'所属组织'" json:"role_id"`
}

// 配额表
type Quota struct {
	BaseModel
	IsShare         int `gorm:"column:is_share;type:int(10);comment:'是否为共享 1 共享 0 独享'" json:"is_share"`
	ResourceGroupID int `gorm:"column:resources_group_id;type:int(10);comment:'资源组ID'" json:"resources_group_id"`
	GroupID         int `gorm:"column:group_id;type:int(10);comment:'组织ID'" json:"group_id"`
	Gpu             int `gorm:"column:gpu;type:int(10);comment:'gpu(显卡) 卡数'" json:"gpu"`
	Cpu             int `gorm:"column:cpu;type:int(10);comment:'cpu 核数'" json:"cpu"`
	Memory          int `gorm:"column:memory;type:int(10);comment:'内存 单位: MB'" json:"memory"`
}

type Group struct {
	BaseModel
	Name          string `gorm:"column:name;type:varchar(50);comment:'组织名称'" json:"name"`
	Module        int    `gorm:"column:module;type:int(10);comment:'所属模块'" json:"module"`
	ParentID      int    `gorm:"column:parent_id;type:int(10);comment:'父级组织ID'" json:"parent_id"`
	LevelPath     int    `gorm:"column:level_path;type:varchar(255);comment:'组织等级路径'" json:"level_path"`
	DiskQuotaSize int    `gorm:"column:disk_quota_size;type:int(10);comment:'磁盘配额大小 单位:TB'" json:"disk_quota_size"`
}

// CREATE TABLE `user` (
//  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
//  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
//  `login_name` varchar(100) NOT NULL DEFAULT '' COMMENT '登录名',
//  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
//  `group_id` int NOT NULL DEFAULT '0' COMMENT '所属组织ID',
//  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
//  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色ID',
// `is_delete` int NOT NULL DEFAULT '0' COMMENT '是否删除，1：删除，0：非删除',
//  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
//  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
//  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
//  `user_type` int DEFAULT NULL COMMENT '1：超级管理员；2：普通用户；3：组管理员;',
//  PRIMARY KEY (`id`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

// create table quota (
//qid serial unique not null,
//is_share int default 0 not null,
//resources_group_id int default null,
//group_id int not null,
//gpu int default 0 not null,
//cpu int default 0 not null,
//memory int default 0 not null,
//#create_time timestamp default current_timestamp not null
//);
//
//comment on table quota is '配额表';
//comment on column quota.qid is '配额ID';
//comment on column quota.is_share is '是否为共享 1 共享 0 独享';
//comment on column quota.resources_group_id is '资源组ID';
//comment on column quota.group_id is '组织ID';
//comment on column quota.gpu is 'gpu(显卡) 卡数';
//comment on column quota.cpu is 'cpu 核数';
//comment on column quota.memory is '内存 单位: MB';

// create table group (
//id serial unique not null,
//name varchar(32) unique not null,
//description varchar(255) default null,
//parent_id int default null,
//level_path varchar(255) not null,
//disk_quota_size int default 0 not null,
//is_delete int default 0 not null,
//#update_time timestamp default current_timestamp,
//#delete_time timestamp default current_timestamp,
//#create_time timestamp default current_timestamp
//
//);
//
//comment on table group is '组织信息表';
//comment on column group.id is '组织ID';
//comment on column group.name is '组织名称';
//comment on column group.description is '组织描述';
//comment on column group.parent_id is '父级组织ID';
//comment on column group.level_path is '组织等级路径';
//comment on column group.disk_quota_size is '磁盘配额大小 单位:TB';
