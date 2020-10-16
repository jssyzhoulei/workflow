package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func groupAPIRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/group", api.GetGroupApis().GroupAddAPI)
	g.GET("/group/quota", api.GetGroupApis().GroupQueryWithQuota)
}