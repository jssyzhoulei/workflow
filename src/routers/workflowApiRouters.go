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
	g.POST("/worknode", api.AddWorkNodes)
}
