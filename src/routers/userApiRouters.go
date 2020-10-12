package routers

import (
	"gitee.com/grandeep/org-svc/client"
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func userApiRouters(g *gin.RouterGroup, o *client.OrgServiceClient) {
	api := apis.NewApis(o.GetUserService())
	g.POST("/user", api.GetUserApis().AddUserApi)
}
