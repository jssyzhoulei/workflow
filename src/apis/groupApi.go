package apis

import (
	"context"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type groupApiI interface {
	GroupAddApi(ctx *gin.Context)
}

type groupApi struct {
	groupService services.GroupServiceI
}

func NewGroupApi(groupService services.GroupServiceI) groupApiI {
	return &groupApi{
		groupService: groupService,
	}
}

// response 通用响应
func response(c *gin.Context, status int, message string, data interface{}) {
	if data == nil {
		data = ""
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(status, map[string]interface{} {
		"code": status,
		"message": message,
		"data": data,
	})
	c.Abort()
	return
}

// GroupAddApi 添加组API
func (u *groupApi) GroupAddApi(c *gin.Context) {

	var data = new(pb_user_v1.GroupAddRequest)

	err := c.BindJSON(data)
	if err != nil {
		response(c, http.StatusBadRequest, "错误", nil)
		return
	}
	_, err = u.groupService.GroupAdd(context.Background(),data)
	if err != nil {
		log.Logger().Error("添加组错误: " + err.Error())
		response(c, http.StatusBadRequest, "错误", nil)
		return
	}
	response(c, http.StatusOK, "成功", nil)
}



