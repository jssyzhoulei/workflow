package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func permissionApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/permission", api.GetPermission().AddPermissionApi)
	g.GET("/menu_cascade", api.GetPermission().GetMenuCascade)
	g.POST("/menu", api.GetPermission().AddMenuApi)
	g.GET("/permission/:id", api.GetPermission().GetPermissionByID)
	g.DELETE("/permission/:id", api.GetPermission().DeletePermissionByID)
	g.PUT("/permission", api.GetPermission().UpdatePermission)
}
