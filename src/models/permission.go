package models

//权限
type Permission struct {
	BaseModel
	UriName    string        `json:"back_end_name" gorm:"column:back_end_name;comment:'后端接口名称';type:varchar(128)"`
	Method     RequestMethod `json:"method" gorm:"column:method;comment:'后端请求method';type:int(2)"`
	Uri        string        `json:"uri" gorm:"column:uri;comment:'后端接口';type:varchar(600)"`
	Relation   int           `json:"relation" gorm:"column:relation;comment:'是否关联前段button 1 是 2 否';type:int(2);default:2"`
	ButtonName string        `json:"button_name" gorm:"column:button_name;comment:'button name';type:varchar(40);"`
	ButtonKey  string        `json:"button_key" gorm:"column:button_key;comment:'前段button key 需要前段全局唯一';type:varchar(40);"`
	MenuID     int64         `json:"menu_id" gorm:"column:menu_id;comment:'菜单id';type:int(11);"`
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
	Version      int `gorm:"column:version;type:int(10);comment:'版本号'" json:"version"`
	RoleID       int `gorm:"column:role_id;type:int(10);comment:'角色id'" json:"role_id"`
	MenuID       int `gorm:"column:component_id;type:int(10);comment:'组件id'" json:"component_id"`
	PermissionID int
}
