package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func groupApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/group", api.GetGroupApis().GroupAddApi)
}