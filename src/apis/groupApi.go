package apis

import (
	"bytes"
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type groupAPIInterface interface {
	GroupAddAPI(ctx *gin.Context)
	GroupQueryWithQuotaAPI(c *gin.Context)
	GroupUpdateAPI(c *gin.Context)
	QuotaUpdateAPI(c *gin.Context)
	GroupTreeQueryAPI(c *gin.Context)
	GroupDelete(c *gin.Context)
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

// GroupAddAPI 添加组API
func (g *groupAPI) GroupAddAPI(c *gin.Context) {

	var data = new(pb_user_v1.GroupAddRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("GroupAdd 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}

	if data.Name == "" || len(data.Quotas) == 0 {
		log.Logger().Info(fmt.Sprintf("GroupAdd 必传参数缺失: name: %d quotas: %v", data.Name, data.Quotas))
		response(c, http.StatusBadRequest, "参数不合法", nil, false)
		return
	}

	l := len(data.Quotas)
	for i := 0; i < l; i++ {
		t := data.Quotas[i]
		if t.IsShare == 0 || strings.Trim(t.ResourcesGroupId, " ") == "" {
			log.Logger().Info(fmt.Sprintf("GroupAdd 必传参数缺失: is_share: %d resources_group_id: %s", t.IsShare,
				t.ResourcesGroupId))
			response(c, http.StatusBadRequest, "参数不合法", nil, false)
			return
		}
	}

	res, err := g.groupService.GroupAddSvc(context.Background(), data)
	if err != nil {
		log.Logger().Info("添加组错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	response(c, http.StatusOK, "成功", res, false)
	return
}

// GroupQueryWithQuotaAPI 查询组和其配额信息
func (g *groupAPI) GroupQueryWithQuotaAPI(c *gin.Context) {
	var data = new(pb_user_v1.GroupQueryWithQuotaByConditionRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("GroupQueryWithQuota 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}

	res, err := g.groupService.GroupQueryWithQuotaByConditionSvc(context.Background(), data)
	if err != nil {
		log.Logger().Info("查询组和其配额信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}

	var _buffer bytes.Buffer

	err = jsonpbMarshaler.Marshal(&_buffer, res)
	if err != nil {
		log.Logger().Info("序列化查询组和其配额信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}

	response(c, http.StatusOK, "成功", _buffer.Bytes(), true)
	return
}

// GroupUpdateAPI 更新组信息
func (g *groupAPI) GroupUpdateAPI(c *gin.Context) {

	var data = new(models.GroupUpdateRequest)
	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("GroupUpdateAPI 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}
	var parentID int64
	var useParentID bool
	if data.ParentID == nil {
		parentID = 0
		useParentID = false
	} else {
		parentID = *data.ParentID
		useParentID = true
	}

	d := &pb_user_v1.GroupUpdateRequest{
		Id:          data.ID,
		Name:        data.Name,
		ParentId:    parentID,
		UseParentId: useParentID,
	}
	resp, err := g.groupService.GroupUpdateSvc(context.Background(), d)
	if err != nil {
		log.Logger().Info("更新组信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	if resp.Code != 0 {
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	response(c, http.StatusOK, "成功", nil, false)
	return
}

// QuotaUpdateAPI 配额更新
func (g *groupAPI) QuotaUpdateAPI(c *gin.Context) {
	var data = new(models.QuotaUpdateRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("QuotaUpdateAPI 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}

	if data.GroupID == 0 || data.ResourcesID == "" || data.IsShare == 0 || data.QuotaType == 0 {
		log.Logger().Info("检测到空值: group_id, is_share, resources_id, quota_type 全部为必传参数")
		response(c, http.StatusBadRequest, "参数错误", nil, false)
		return
	}

	d := &pb_user_v1.QuotaUpdateRequest{
		GroupId:     data.GroupID,
		IsShare:     data.IsShare,
		ResourcesId: data.ResourcesID,
		QuotaType:   data.QuotaType,
		Total:       data.Total,
		Used:        data.Used,
	}

	resp, err := g.groupService.QuotaUpdateSvc(context.Background(), d)
	if err != nil {
		log.Logger().Info("更新配额信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	if resp.Code != 0 {
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	response(c, http.StatusOK, "成功", nil, false)
	return
}

// GroupTreeQueryAPI 组树查询
func (g *groupAPI) GroupTreeQueryAPI(c *gin.Context) {
	groupIDStr := c.DefaultQuery("group_id", "")
	if groupIDStr == "" {
		response(c, http.StatusBadRequest, "获取参数失败", nil, false)
		return
	}

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		response(c, http.StatusBadRequest, "参数解析失败", nil, false)
		return
	}

	d := &pb_user_v1.GroupID{
		Id: groupID,
	}

	resp, err := g.groupService.GroupTreeQuerySvc(context.Background(), d)
	if err != nil {
		log.Logger().Info("获取 组树失败: " + err.Error())
		response(c, http.StatusBadRequest, "查询失败或结果为空", nil, false)
		return
	}
	response(c, http.StatusOK, "成功", resp.TreeJson, true)
	return
}

// GroupDelete 删除组
func (g *groupAPI) GroupDelete(c *gin.Context) {
	groupIDStr := c.DefaultQuery("group_id", "")
	if groupIDStr == "" {
		response(c, http.StatusBadRequest, "获取参数失败", nil, false)
		return
	}

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		response(c, http.StatusBadRequest, "参数解析失败", nil, false)
		return
	}

	d := &pb_user_v1.GroupID{
		Id: groupID,
	}

	_, err = g.groupService.GroupDeleteSvc(context.Background(), d)
	if err != nil {
		log.Logger().Info("删除 组失败: " + err.Error())
		response(c, http.StatusBadRequest, "删除失败", nil, false)
		return
	}
	response(c, http.StatusOK, "成功", nil, false)
	return

}
