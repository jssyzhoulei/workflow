package apis

import (
	//"gitee.com/grandeep/org-svc/logger"
	"github.com/gin-gonic/gin"
)

type groupAPIInterface interface {
	GroupAddAPI(ctx *gin.Context)
	GroupQueryWithQuotaAPI(c *gin.Context)
	GroupUpdateAPI(c *gin.Context)
	QuotaUpdateAPI(c *gin.Context)
	GroupTreeQueryAPI(c *gin.Context)
	GroupDelete(c *gin.Context)
	QueryGroupAndSubGroupsUsers(c *gin.Context)
	SetGroupQuotaUsed(c *gin.Context)
	QuerySubGroupsUsers(c *gin.Context)
	QueryGroupIDAndSubGroupsID(c *gin.Context)
	QueryTopGroupExcludeSelfUsers(c *gin.Context)
}

