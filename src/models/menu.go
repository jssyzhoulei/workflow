package models

type Menu struct {
	BaseModel
	Name         string     `gorm:"column:name;type:varchar(128);not null;unique_index;comment:'菜单名称'" json:"name"`
	ParentID     int        `gorm:"column:parent_id;type:int(10);comment:'父组件'" json:"parent_id"`
	Module       MenuModule `gorm:"column:module;type:int(10);comment:'所属模块'" json:"module"`
	Order        int        `gorm:"column:order;type:int(10);comment:'组件次序'" json:"order"`
	Version      int        `gorm:"column:version;type:int(10);comment:'版本号'" json:"version"`
	TemplatePath string     `gorm:"column:template_path;type:varchar(128);comment:'前段组件路径'" json:"template_path"`
	Status       int        `gorm:"column:status;type:int(2);not null;comment:'菜单状态 1 启用 2 未启用';default:1" json:"status"`
}

func (m Menu) TableName() string {
	return "menu"
}

//业务模块
type MenuModule int8

const (
	MODULE_TOP MenuModule = iota
	// 基础模块 - 算力云
	MODULE_BASIC
	// 标注模块
	MODULE_ANNOTATION
	MODULE_TRAINING
	MODULE_DEVELOP
	MODULE_SERVICE
)

type Cascade struct {
	Value int `json:"value"`
	Label string `json:"label"`
	Children []Cascade `json:"children"`
}


func (m Menu) GetMenuCascade(menus []Menu, parentId int) (cascades []Cascade) {
	for k, menu := range menus {
		var (
			cascade Cascade
		)
		if parentId == menu.ParentID {
			var (
				menusNew = make([]Menu, len(menus)-1)
			)
			copy(menusNew[:k], menus[:k])
			copy(menusNew[k:], menus[k+1:])
			cascade.Value = menu.ID
			cascade.Label = menu.Name
			cascade.Children = m.GetMenuCascade(menusNew, menu.ID)
			cascades = append(cascades, cascade)
		}
	}
	return cascades
}

func (m Menu) AddPermissionCascade(permissions []Permission, cascades []Cascade) []Cascade {
	for k, cascade := range cascades {
		if len(cascade.Children) > 0 {
			cs := m.AddPermissionCascade(permissions, cascade.Children)
			cascades[k].Children = cs
			continue
		} else {
			for index, permission := range permissions {
				if cascade.Value == permission.MenuID {
					var (
						c Cascade
						permissionsNew = make([]Permission, len(permissions) -1)
					)
					copy(permissionsNew[:index], permissions[:index])
					copy(permissionsNew[index:], permissions[index+1:])
					c.Value = permission.ID
					c.Label = permission.UriName
					cascade.Children = append(cascade.Children, c)
				}
			}
			cascades[k] = cascade
		}
	}
	return cascades
}
