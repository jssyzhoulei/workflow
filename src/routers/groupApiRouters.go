package routers

import (
	"gitee.com/grandeep/org-svc/src/apis"
	"github.com/gin-gonic/gin"
)

func groupAPIRouters(g *gin.RouterGroup, api apis.IApis) {
	g.POST("/group", api.GetGroupApis().GroupAddAPI)
	g.PUT("/group", api.GetGroupApis().GroupUpdateAPI)
	g.DELETE("/group", api.GetGroupApis().GroupDelete)
	g.POST("/group/quota", api.GetGroupApis().GroupQueryWithQuotaAPI)
	g.GET("/group/tree", api.GetGroupApis().GroupTreeQueryAPI)
	g.GET("/group/user", api.GetGroupApis().QueryGroupAndSubGroupsUsers)
	g.GET("/group/sub_user", api.GetGroupApis().QuerySubGroupsUsers)
	g.GET("/group/sub_group_id", api.GetGroupApis().QueryGroupIDAndSubGroupsID)
	g.PUT("/quota", api.GetGroupApis().QuotaUpdateAPI)
	g.PUT("/quota/used", api.GetGroupApis().SetGroupQuotaUsed)
}