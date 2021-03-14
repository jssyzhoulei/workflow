package services
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"gitee.com/grandeep/org-svc/cmd/org-svc/engine"
//	"gitee.com/grandeep/org-svc/src/models"
//	"gitee.com/grandeep/org-svc/src/repositories"
//	"testing"
//)
//var (
//	background = context.Background()
//	configPath = "/Users/fanhaojie/dew/app/Go/Fu/kit-fu/org-svc/resources/config/config.yaml"
//	e = engine.NewEngine(configPath)
//	repo = repositories.NewRepoI(e.DB)
//)
//func Test_permissionService_AddMenu(t *testing.T) {
//	var (
//		menu = models.Menu{
//			Name:         "首页父2子2",
//			ParentID:     9,
//			Module:      models.MODULE_BASIC,
//			Order:        0,
//			Version:      0,
//			TemplatePath: "/index",
//			Status:       0,
//		}
//	)
//
//	fmt.Println(NewPermissionService(repo).AddMenuSvc(background, menu))
//}
//
//func Test_permissionService_UpdateMenuByID(t *testing.T) {
//	var (
//		menu = models.Menu{
//			BaseModel: models.BaseModel{
//				ID:            6,
//			},
//			Name:         "首页woca",
//			ParentID:     0,
//			Module:      models.MODULE_BASIC,
//			Order:        0,
//			Version:      0,
//			TemplatePath: "/index",
//			Status:       0,
//		}
//	)
//	fmt.Println(NewPermissionService(repo).UpdateMenuByIDSvc(background, menu))
//}
//
//func Test_permissionService_AddPermission(t *testing.T) {
//	permission := models.Permission{
//		UriName:    "首页数据接",
//		Method:     models.METHOD_GET,
//		Uri:        "/apis/v1/index",
//		Relation:   1,
//		ButtonName: "首页",
//		ButtonKey:  "index",
//		MenuID:     1,
//	}
//	NewPermissionService(repo).AddPermissionSvc(background, permission)
//}
//
//func Test_permissionService_GetPermissionByID(t *testing.T) {
//	fmt.Println(NewPermissionService(repo).GetPermissionByIDSvc(background, 1))
//}
//
//func Test_permissionService_UpdatePermissionByID(t *testing.T) {
//	permission := models.Permission{
//		BaseModel:models.BaseModel{ID: 1},
//		UriName:    "首页数据接",
//		Method:     models.METHOD_GET,
//		Uri:        "/apis/v1/index",
//		Relation:   1,
//		ButtonName: "首页",
//		ButtonKey:  "index",
//		MenuID:     1,
//	}
//	NewPermissionService(repo).UpdatePermissionByIDSvc(background, permission)
//}
//
//func Test_permissionService_DeletePermissionByID(t *testing.T) {
//	fmt.Println(NewPermissionService(repo).DeletePermissionByIDSvc(background, 1))
//}
//
//func Test_permissionService_GetMenuCascadeByModule(t *testing.T) {
//	//fmt.Println(NewPermissionService(repo).GetMenuCascadeByModuleSvc(background, 1))
//	rs := makePermissionRouteTree([]models.Permission{{
//		Uri:"/apis/v1/user/:id/a/:id/b",
//		Method: models.METHOD_GET,
//	},{
//		Uri:"/apis/v1/user/a",
//		Method: models.METHOD_GET,
//	},{
//		Uri:"/apis/v1/user/",
//		Method: models.METHOD_DELETE,
//	}})
//	fmt.Println(rs.AuthRoute("/apis/v1/user//DELETE"))
//	fmt.Println(rs)
//}
//
//type routeTree map[string]routeTree
//
//func (rt routeTree) AuthRoute(uri string) error {
//	var route []byte
//	var err error
//	for i:= 0; i<len(uri); i++ {
//		switch uri[i] {
//		case '/':
//			rs := string(route)
//			if len(route) > 0 {
//				if _, ok := rt[rs]; ok {
//					err = rt[rs].AuthRoute(uri[i+1:])
//					if err != nil {
//						if _, ok := rt["*"]; ok {
//							err = rt["*"].AuthRoute(uri[i+1:])
//						}
//					}
//					return err
//				}
//				if _, ok := rt["*"]; ok {
//					err = rt["*"].AuthRoute(uri[i+1:])
//					return err
//				}
//				return errors.New("no permission")
//			} else {
//				return rt.AuthRoute(uri[i+1:])
//			}
//		default:
//			route = append(route, uri[i])
//		}
//	}
//
//	if route != nil {
//		if _, ok := rt["*"]; ok {
//			return nil
//		}
//		fmt.Println(string(route))
//		if _, ok := rt[string(route)]; ok {
//			return nil
//		}
//	}
//	return errors.New("no permission")
//}
//
//func makePermissionRouteTree(ps []models.Permission) routeTree {
//	rt := make(routeTree)
//	for _, v := range ps {
//		makeUri(v.Uri + "/" + v.Method.String(), rt)
//	}
//	return rt
//}
//
//func makeUri(uri string, rt routeTree) {
//	var b bool
//	var route []byte
//	for i:= 0; i<len(uri); i++ {
//		switch uri[i] {
//		case '/':
//			rs := string(route)
//			if len(route) > 0 {
//				if b {
//					if _, ok := rt["*"]; !ok {
//						rt["*"] = make(routeTree)
//					}
//					makeUri(uri[i+1:], rt["*"])
//				} else {
//					if _, ok := rt[rs]; !ok {
//						rt[rs] = make(routeTree)
//					}
//					makeUri(uri[i+1:], rt[rs])
//				}
//			} else {
//				makeUri(uri[i+1:], rt)
//			}
//			return
//		case '*':
//			b = true
//		case ':':
//			b = true
//		default:
//			route = append(route, uri[i])
//		}
//	}
//	if route != nil {
//		if b {
//			rt["*"] = make(routeTree)
//		} else {
//			rt[string(route)] = make(routeTree)
//		}
//	}
//}
