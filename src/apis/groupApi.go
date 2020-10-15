package apis

import (
	"context"
	"fmt"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type groupAPIInterface interface {
	GroupAddAPI(ctx *gin.Context)
	GroupQueryWithQuota(c *gin.Context)
}

type groupAPI struct {
	groupService services.GroupServiceInterface
}

// NewGroupAPI ...
func NewGroupAPI(groupService services.GroupServiceInterface) groupAPIInterface {
	return &groupAPI{
		groupService: groupService,
	}
}

// response 通用响应
func response(c *gin.Context, status int, message string, data interface{}) {
	if data == nil {
		data = ""
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    status,
		"message": message,
		"data":    data,
	})
	c.Abort()
	return
}

// GroupAddApi 添加组API
func (g *groupAPI) GroupAddAPI(c *gin.Context) {

	var data = new(pb_user_v1.GroupAddRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Warn(fmt.Sprintf("GroupAdd 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil)
		return
	}
	_, err = g.groupService.GroupAddSvc(context.Background(), data)
	if err != nil {
		log.Logger().Error("添加组错误: " + err.Error())
		response(c, http.StatusBadRequest, "错误", nil)

		if data.Name == "" || len(data.Quotas) == 0 {
			log.Logger().Warn(fmt.Sprintf("GroupAdd 必传参数缺失: name: %d quotas: %v", data.Name, data.Quotas))
			response(c, http.StatusBadRequest, "参数不合法", nil)
			return
		}

		l := len(data.Quotas)
		for i := 0; i < l; i++ {
			t := data.Quotas[i]
			if t.IsShare == 0 || strings.Trim(t.ResourcesGroupId, " ") == "" {
				log.Logger().Warn(fmt.Sprintf("GroupAdd 必传参数缺失: is_share: %d resources_group_id: %s", t.IsShare,
					t.ResourcesGroupId))
				response(c, http.StatusBadRequest, "参数不合法", nil)
				return
			}
		}

		res, err := g.groupService.GroupAddSvc(context.Background(), data)
		if err != nil {
			log.Logger().Warn("添加组错误: " + err.Error())
			response(c, http.StatusBadRequest, "操作失败", nil)
			return
		}

		response(c, http.StatusOK, "成功", res)
		return
	}
}

// GroupQueryWithQuota 查询组和其配额信息
func(g *groupAPI) GroupQueryWithQuota(c * gin.Context){
	var data = new(pb_user_v1.GroupQueryWithQuotaByConditionRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Warn(fmt.Sprintf("GroupQueryWithQuota 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil)
		return
	}

	res, err := g.groupService.GroupQueryWithQuotaByConditionSvc(context.Background(), data)
	if err != nil {
		log.Logger().Warn("查询组和其配额信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil)
		return
	}

	response(c, http.StatusOK, "成功", res.Groups)
	return
}