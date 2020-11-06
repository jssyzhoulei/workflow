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
	g.POST("/user/list", api.GetUserApis().GetUsersApi)
	g.POST("/user/import_user", api.GetUserApis().ImportUser)
	g.POST("/user/batch_del",api.GetUserApis().BatchDeleteUsersApi)
	g.PUT("/user/by_group", api.GetUserApis().ImportUsersByGroupIdApi)
	g.GET("/user/import_template.xlsx", api.GetUserApis().ImportUserTemplate)
}
