package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func userApiRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/user", api.GetUserApis().AddUserApi)
	g.PUT("/user", api.GetUserApis().UpdateUserByIDApi)
	g.GET("/user/:id", api.GetUserApis().GetUserByIDApi)
	g.DELETE("/user/:id", api.GetUserApis().DeleteUserByIDApi)
	g.POST("/user/import_user", api.GetUserApis().ImportUser)
}
