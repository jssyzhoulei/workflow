package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func groupAPIRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/group", api.GetGroupApis().GroupAddAPI)
	g.PUT("/group", api.GetGroupApis().GroupUpdateAPI)
	g.DELETE("/group", api.GetGroupApis().GroupDelete)
	g.GET("/group/quota", api.GetGroupApis().GroupQueryWithQuotaAPI)
	g.GET("/group/tree", api.GetGroupApis().GroupTreeQueryAPI)
	g.GET("/group/user", api.GetGroupApis().QueryGroupAndSubGroupsUsers)
	g.PUT("/quota", api.GetGroupApis().QuotaUpdateAPI)
}