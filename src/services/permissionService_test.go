package services

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/utils/src/pkg/engine"
	"testing"
)
var (
	background = context.Background()
	configPath = "/Users/fanhaojie/dew/app/Go/Fu/kit-fu/org-svc/resources/config/config.yaml"
)
func Test_permissionService_AddMenu(t *testing.T) {
	var (
		menu = models.Menu{
			Name:         "首页父2子2",
			ParentID:     9,
			Module:      models.MODULE_BASIC,
			Order:        0,
			Version:      0,
			TemplatePath: "/index",
			Status:       0,
		}
	)
	e := engine.NewEngine(configPath)
	fmt.Println(NewPermissionService(e).AddMenuSvc(background, menu))
}

func Test_permissionService_UpdateMenuByID(t *testing.T) {
	var (
		menu = models.Menu{
			BaseModel: models.BaseModel{
				ID:            6,
			},
			Name:         "首页woca",
			ParentID:     0,
			Module:      models.MODULE_BASIC,
			Order:        0,
			Version:      0,
			TemplatePath: "/index",
			Status:       0,
		}
	)
	e := engine.NewEngine(configPath)
	fmt.Println(NewPermissionService(e).UpdateMenuByIDSvc(background, menu))
}

func Test_permissionService_AddPermission(t *testing.T) {
	permission := models.Permission{
		UriName:    "首页数据接",
		Method:     models.METHOD_GET,
		Uri:        "/apis/v1/index",
		Relation:   1,
		ButtonName: "首页",
		ButtonKey:  "index",
		MenuID:     1,
	}
	e := engine.NewEngine(configPath)
	NewPermissionService(e).AddPermissionSvc(background, permission)
}

func Test_permissionService_GetPermissionByID(t *testing.T) {
	e := engine.NewEngine(configPath)
	fmt.Println(NewPermissionService(e).GetPermissionByIDSvc(background, 1))
}

func Test_permissionService_UpdatePermissionByID(t *testing.T) {
	permission := models.Permission{
		BaseModel:models.BaseModel{ID: 1},
		UriName:    "首页数据接",
		Method:     models.METHOD_GET,
		Uri:        "/apis/v1/index",
		Relation:   1,
		ButtonName: "首页",
		ButtonKey:  "index",
		MenuID:     1,
	}
	e := engine.NewEngine(configPath)
	NewPermissionService(e).UpdatePermissionByIDSvc(background, permission)
}

func Test_permissionService_DeletePermissionByID(t *testing.T) {
	e := engine.NewEngine(configPath)
	fmt.Println(NewPermissionService(e).DeletePermissionByIDSvc(background, 1))
}

func Test_permissionService_GetMenuCascadeByModule(t *testing.T) {
	e := engine.NewEngine(configPath)
	fmt.Println(NewPermissionService(e).GetMenuCascadeByModuleSvc(background, 1))
}