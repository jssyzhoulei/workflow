package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jssyzhoulei/workflow/src/apis"
)

func workApiRouters(g *gin.RouterGroup, api *apis.Apis) {
	g.POST("/workflow", api.AddWorkflow)
	g.PUT("/workflow", api.UpdateWorkflow)
	g.DELETE("/workflow", api.DelWorkflow)
	g.GET("/workflow", api.ListWorkflow)
	//g.PUT("/user", api.GetUserApis().UpdateUserByIDApi)
	//g.GET("/user/query/:id", api.GetUserApis().GetUserByIDApi)
	//g.DELETE("/user/del/:id", api.GetUserApis().DeleteUserByIDApi)
	//g.POST("/user/list", api.GetUserApis().GetUsersApi)
	//g.POST("/user/import_user", api.GetUserApis().ImportUser)
	//g.POST("/user/batch_del",api.GetUserApis().BatchDeleteUsersApi)
	//g.PUT("/user/by_group", api.GetUserApis().ImportUsersByGroupIdApi)
	//g.GET("/user/import_template.xlsx", api.GetUserApis().ImportUserTemplate)
}
