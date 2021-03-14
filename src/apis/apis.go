package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jssyzhoulei/workflow/src/apis/code"
	"net/http"
)


type IApis interface {
	GetUserApis() userApiInterface
	GetPermission() permissionApiInterface
	GetGroupApis() groupAPIInterface
	GetRoleApis() RoleApiInterface
}

type apis struct {
	userApiInterface
	permissionApiInterface permissionApiInterface
	groupAPIInterface
	RoleApiInterface
}

func (a *apis) GetUserApis() userApiInterface {
	return a.userApiInterface
}

func (a *apis) GetGroupApis() groupAPIInterface {
	return a.groupAPIInterface
}

func (a *apis) GetRoleApis() RoleApiInterface {
	return a.RoleApiInterface
}

func (a *apis) GetPermission() permissionApiInterface {
	return a.permissionApiInterface
}

func success_(c *gin.Context, data interface{}) {
	if data == nil {
		data = ""
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(200, ApiResponse{
		Code:    code.OK,
		Message: "ok",
		Data:    data,
	})
	c.Abort()
	return
}

func error_(c *gin.Context, status code.Code, err ...error) {
	//c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(200, ApiResponse{
		Code:    status,
		Message: status.Message(err...),
		Data:    nil,
	})
	c.Abort()
	return
}

// response 通用响应
// @data 当 isPB(是否返回的是 jsonpb 处理的数据) 为 true 时, data 必须为 []byte 参考 apis.groupAPIInterface GroupQueryWithQuotaAPI 方法
func response(c *gin.Context, status int, message string, data interface{}, isByte bool) {
	if data == nil {
		data = ""
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	if !isByte {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": message,
			"data":    data,
		})
	} else {
		_, err := c.Writer.Write(data.([]byte))
		if err != nil {
			log.Logger.Warn("PB消息byte写入响应信息失败: " + err.Error())
		}
	}

	c.Abort()
	return
}

//var _buffer bytes.Buffer
//
//err = jsonpbMarshaler.Marshal(&_buffer, res)
//if err != nil {
//	log.Logger().Info("序列化查询组和其配额信息错误: " + err.Error())
//	response(c, http.StatusBadRequest, "操作失败", nil, false)
//	return
//}

type ApiResponse struct {
	Code    code.Code   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
