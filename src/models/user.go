package models

type User2 struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

// User 用户表
func (u User2) TableName() string {
	return "user2"
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

// TableName ...
func (u User) TableName() string {
	return "user"
}

// UserRole 用户角色关联表
type UserRole struct {
	BaseModel
	UserID int `gorm:"column:user_id;type:int(10);comment:'关联用户'" json:"user_id"`
	RoleID int `gorm:"column:role_id;type:int(10);comment:'关联角色'" json:"role_id"`
}

// TableName ...
func (u UserRole) TableName() string {
	return "user_role"
}

//导入用户请求
type ImportUserRequest struct {
	RoleID []int64 `json:"role_id"`
	GroupID int64 `json:"group_id"`
	Content string `json:"content"`
	IsCover int32 `json:"is_cover"`
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
