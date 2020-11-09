package apis

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
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
		log.Logger().Info(fmt.Sprintf("GroupAdd 必传参数缺失: name: %s quotas: %v", data.Name, data.Quotas))
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

	var result = make([]*models.GroupQueryWithQuota, 0)
	for i := 0; i < len(res.Groups); i++ {
		_data := res.Groups[i]

		var quotaResult = make([]*models.QuotaResponse, 0)

		for i := 0; i < len(_data.Quotas); i++ {
			q := _data.Quotas[i]
			_tmp := &models.QuotaResponse{
				IsShare:          q.IsShare,
				ResourcesGroupId: q.ResourcesGroupId,
				Cpu:              q.Cpu,
				Gpu:              q.Gpu,
				Memory:           q.Memory,
			}
			quotaResult = append(quotaResult, _tmp)
		}

		var _tmp = &models.GroupQueryWithQuota{
			ID:            _data.Id,
			Name:          _data.Name,
			ParentID:      _data.ParentId,
			TopParentID:   _data.TopParentId,
			DiskQuotaSize: _data.DiskQuotaSize,
			Description:   _data.Description,
			Quotas:        quotaResult,
		}
		result = append(result, _tmp)
	}

	response(c, http.StatusOK, "成功", result, false)
	return
}

// GroupUpdateAPI 更新组信息
func (g *groupAPI) GroupUpdateAPI(c *gin.Context) {

	var data = new(models.GroupUpdateRequest)
	err := c.BindJSON(&data)
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

	var pbQuotas = make([]*pb_user_v1.Quota, 0)
	for _, v := range data.Quotas {
		_t := &pb_user_v1.Quota{
			IsShare:          v.IsShare,
			ResourcesGroupId: v.ResourcesGroupId,
			Gpu:              v.Gpu,
			Cpu:              v.Cpu,
			Memory:           v.Memory,
		}
		pbQuotas = append(pbQuotas, _t)
	}

	d := &pb_user_v1.GroupUpdateRequest{
		Id:            data.ID,
		Name:          data.Name,
		ParentId:      parentID,
		UseParentId:   useParentID,
		Description:   data.Description,
		DiskQuotaSize: data.DiskQuotaSize,
		Quotas:        pbQuotas,
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
	response(c, http.StatusOK, "成功", string(resp.TreeJson), false)
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

// QueryGroupAndSubGroupsUsers 查询组及其子组下的所有用户
func (g *groupAPI) QueryGroupAndSubGroupsUsers(c *gin.Context) {
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

	resp, err := g.groupService.QueryGroupAndSubGroupsUsersSvc(context.Background(), d)
	if err != nil {
		log.Logger().Info("查询组及其子组下的所有用户 失败: " + err.Error())
		response(c, http.StatusBadRequest, "查询失败", nil, false)
		return
	}

	var result = make([]*models.QueryGroupsUsersResponse, 0)
	for i := 0; i < len(resp.Users); i++ {
		_user := resp.Users[i]
		_tmp := &models.QueryGroupsUsersResponse{
			ID:        _user.Id.Id,
			UserName:  _user.UserName,
			LoginName: _user.LoginName,
			GroupID:   _user.GroupId,
			UserType:  int(_user.UserType),
			Mobile:    _user.Mobile,
		}
		result = append(result, _tmp)
	}
	response(c, http.StatusOK, "成功", result, false)
	return
}

// SetGroupQuotaUsed 设置组配额已使用
func (g *groupAPI) SetGroupQuotaUsed(c *gin.Context) {

	var data = &models.SetGroupQuotaRequest{}

	err := c.BindJSON(&data)
	if err != nil {
		log.Logger().Info("SetGroupQuotaUsed 解析参数失败: " + err.Error())
		response(c, http.StatusBadRequest, "解析参数失败", nil, false)
		return
	}

	var d = &pb_user_v1.SetGroupQuotaUsedRequest{
		GroupId:   data.GroupID,
		IsShare:   data.IsShare,
		QuotaType: data.QuotaType,
		Used:      data.Used,
	}

	_, err = g.groupService.SetGroupQuotaUsedSvc(context.Background(), d)
	if err != nil {
		log.Logger().Error("SetGroupQuotaUsed 操作失败: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	response(c, http.StatusOK, "成功", nil, false)
	return

}

// QuerySubGroupsUsers 查询子组下的用户
func (g *groupAPI) QuerySubGroupsUsers(c *gin.Context) {
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

	resp, err := g.groupService.QuerySubGroupsUsersSvc(context.Background(), d)
	if err != nil {
		log.Logger().Info("查询子组下的用户 失败: " + err.Error())
		response(c, http.StatusBadRequest, "查询失败", nil, false)
		return
	}

	var result = make([]*models.QueryGroupsUsersResponse, 0)
	for i := 0; i < len(resp.Users); i++ {
		_user := resp.Users[i]
		_tmp := &models.QueryGroupsUsersResponse{
			ID:        _user.Id.Id,
			UserName:  _user.UserName,
			LoginName: _user.LoginName,
			GroupID:   _user.GroupId,
			UserType:  int(_user.UserType),
			Mobile:    _user.Mobile,
		}
		result = append(result, _tmp)
	}
	response(c, http.StatusOK, "成功", result, false)
	return
}

// QueryGroupIDAndSubGroupsID 查询组下的所有子组ID包含被查询的组ID
func (g *groupAPI) QueryGroupIDAndSubGroupsID(c *gin.Context) {

	groupIDStr := c.DefaultQuery("group_id", "")

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		log.Logger().Info("QueryGroupIDAndSubGroupsID 解析参数错误: " + err.Error())
		response(c, http.StatusBadRequest, "解析参数错误", nil, false)
		return
	}

	data := &pb_user_v1.GroupID{
		Id:                   groupID,
	}

	resp, err := g.groupService.QueryGroupIDAndSubGroupsIDSvc(context.Background(), data)
	if err != nil {
		log.Logger().Info("QueryGroupIDAndSubGroupsID 查询错误: " + err.Error())
		response(c, http.StatusBadRequest, "获取用户组子组失败或为空", nil, false)
		return
	}

	response(c, http.StatusOK, "成功", resp.Ids, false)
	return
}
