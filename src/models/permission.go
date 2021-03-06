package models

//权限
type Permission struct {
	BaseModel
	UriName    string        `json:"uri_name" gorm:"column:uri_name;comment:'后端接口名称';type:varchar(128)"`
	Method     RequestMethod `json:"method" gorm:"column:method;comment:'后端请求method';type:int(2)"`
	Uri        string        `json:"uri" gorm:"column:uri;comment:'后端接口';type:varchar(600)"`
	Relation   int           `json:"relation" gorm:"column:relation;comment:'是否关联前段button 1 是 2 否';type:int(2);default:2"`
	ButtonName string        `json:"button_name" gorm:"column:button_name;comment:'button name';type:varchar(40);"`
	ButtonKey  string        `json:"button_key" gorm:"column:button_key;comment:'前段button key 需要前段全局唯一';type:varchar(40);"`
	MenuID     int           `json:"menu_id" gorm:"column:menu_id;comment:'菜单id';type:int(11);"`
	Module     MenuModule    `json:"module" gorm:"column:module; comment:'所属模块平台';type:int(2)" json:"module"`
}

func (p Permission) TableName() string {
	return "permission"
}

type RequestMethod int8

const (
	METHOD_GET = iota + 1
	METHOD_POST
	METHOD_PUT
	METHOD_DELETE
)

func (r RequestMethod) String() string {
	switch r {
	case METHOD_GET:
		return "GET"
	case METHOD_POST:
		return "POST"
	case METHOD_PUT:
		return "PUT"
	case METHOD_DELETE:
		return "DELETE"
	default:
		return ""
	}
}

type RoleMenuPermission struct {
	BaseModel
	Version      int `gorm:"column:version;type:int(10);comment:'版本号'" json:"-"`
	RoleID       int `gorm:"column:role_id;type:int(10);comment:'角色id'" json:"role_id"`
	MenuID       int `gorm:"column:menu_id;type:int(10);comment:'组件id'" json:"menu_id"`
	PermissionID int `json:"permission_id"`
}

func (RoleMenuPermission) TableName() string {
	return "role_menu_permission"
}

type CreateMenuPermRequest struct {
	Role
	MenuPerms []*RoleMenuPermission `json:"menu_perms"`
}

func (c CreateMenuPermRequest) Check() bool {

	return len(c.MenuPerms) != 0 && c.Name != "" && (c.DataPermit == 1 || c.DataPermit == 2 || c.DataPermit == 3)
}

type Perm struct {
	RoleID       int `gorm:"column:role_id;type:int(10);comment:'角色id'" json:"role_id"`
	MenuID       int `gorm:"column:menu_id;type:int(10);comment:'组件id'" json:"menu_id"`
	PermissionID int `json:"permission_id"`
}

type MenuPermResponse struct {
	Role
	Perm
}
