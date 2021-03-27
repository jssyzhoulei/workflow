package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jssyzhoulei/workflow/logger"
	"github.com/jssyzhoulei/workflow/src/apis/code"
	"github.com/jssyzhoulei/workflow/src/services"
	"net/http"
)

type Apis struct {
	*services.WorkService
}

func NewApis(s *services.WorkService) *Apis {
	return &Apis{s}
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
			log.Logger().Warn("PB消息byte写入响应信息失败: " + err.Error())
		}
	}

	c.Abort()
	return
}

type ApiResponse struct {
	Code    code.Code   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
