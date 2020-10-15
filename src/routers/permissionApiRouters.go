package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func permissionApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/permission", api.GetPermission().AddPermissionApi)
}
