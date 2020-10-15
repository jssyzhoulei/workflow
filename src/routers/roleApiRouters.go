package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func roleApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/role", api.GetRoleApis().AddRoleApi)
	g.PUT("/role", api.GetRoleApis().UpdateRoleApi)
	g.DELETE("/role", api.GetRoleApis().DeleteRoleApi)
}
