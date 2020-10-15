package apis

import (
	"bytes"
	"context"
	"fmt"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"net/http"
	"strings"
)

var (
	jsonpbMarshaler *jsonpb.Marshaler
)

func init() {
	jsonpbMarshaler = &jsonpb.Marshaler{
		EnumsAsInts : true,
		EmitDefaults: true,
		OrigName    : true,
	}
}

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
// @data 当 isPB 为 true 时, data 必须为 []byte
func response(c *gin.Context, status int, message string, data interface{}, isPB bool) {
	if data == nil {
		data = ""
	}
	if !isPB {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": message,
			"data":    data,
		})
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
		_, err := c.Writer.Write(data.([]byte))
		if err != nil {
			log.Logger().Warn("PB消息byte写入响应信息失败: " + err.Error())
		}
	}

	c.Abort()
	return
}

// GroupAddApi 添加组API
func (g *groupAPI) GroupAddAPI(c *gin.Context) {

	var data = new(pb_user_v1.GroupAddRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Warn(fmt.Sprintf("GroupAdd 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}

	if data.Name == "" || len(data.Quotas) == 0 {
		log.Logger().Warn(fmt.Sprintf("GroupAdd 必传参数缺失: name: %d quotas: %v", data.Name, data.Quotas))
		response(c, http.StatusBadRequest, "参数不合法", nil, false)
		return
	}

	l := len(data.Quotas)
	for i := 0; i < l; i++ {
		t := data.Quotas[i]
		if t.IsShare == 0 || strings.Trim(t.ResourcesGroupId, " ") == "" {
			log.Logger().Warn(fmt.Sprintf("GroupAdd 必传参数缺失: is_share: %d resources_group_id: %s", t.IsShare,
				t.ResourcesGroupId))
			response(c, http.StatusBadRequest, "参数不合法", nil, false)
			return
		}
	}

	res, err := g.groupService.GroupAddSvc(context.Background(), data)
	if err != nil {
		log.Logger().Warn("添加组错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}

	response(c, http.StatusOK, "成功", res, false)
	return
}

// GroupQueryWithQuota 查询组和其配额信息
func (g *groupAPI) GroupQueryWithQuota(c *gin.Context) {
	var data = new(pb_user_v1.GroupQueryWithQuotaByConditionRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Warn(fmt.Sprintf("GroupQueryWithQuota 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}

	res, err := g.groupService.GroupQueryWithQuotaByConditionSvc(context.Background(), data)
	if err != nil {
		log.Logger().Warn("查询组和其配额信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}

	var _buffer bytes.Buffer

	err = jsonpbMarshaler.Marshal(&_buffer, res)
	if err != nil {
		log.Logger().Warn("序列化查询组和其配额信息错误: " + err.Error())
		response(c, http.StatusBadRequest, "操作失败", nil, false)
		return
	}

	response(c, http.StatusOK, "成功", _buffer.Bytes(), true)
	return
}

