package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func userApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/user", api.GetUserApis().AddUserApi)
}
