package apis

import (
	"gitee.com/grandeep/org-svc/client"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"net/http"
)

var (
	jsonpbMarshaler *jsonpb.Marshaler
)

func init() {
	jsonpbMarshaler = &jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
		OrigName:     true,
	}
}

type IApis interface {
	GetUserApis() userApiI
	GetGroupApis() groupAPIInterface

}

type apis struct {
	userApiI
	groupAPIInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		groupAPIInterface: NewGroupAPI(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetGroupApis() groupAPIInterface {
	return a.groupAPIInterface
}


// response 通用响应
// @data 当 isPB(是否返回的是 jsonpb 处理的数据) 为 true 时, data 必须为 []byte 参考 apis.groupAPIInterface GroupQueryWithQuotaAPI 方法
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