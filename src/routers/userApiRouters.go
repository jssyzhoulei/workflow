package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func userApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/user", api.GetUserApis().AddUserApi)
	g.PUT("/user", api.GetUserApis().UpdateUserByIDApi)
	g.GET("/user/query/:id", api.GetUserApis().GetUserByIDApi)
	g.DELETE("/user/del/:id", api.GetUserApis().DeleteUserByIDApi)
	g.GET("/user/list", api.GetUserApis().GetUserListApi)
}
