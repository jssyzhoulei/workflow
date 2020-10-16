package apis

import (
	"gitee.com/grandeep/org-svc/client"
	"gitee.com/grandeep/org-svc/src/apis/code"
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

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiInterface: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
		//groupApiI,NewGroupApi(o.),
		RoleApiInterface:NewRoleApi(o.GetRoleService()),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
		groupAPIInterface: NewGroupAPI(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiInterface {
	return a.userApiInterface
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
	c.JSON(200, ApiResponse {
		Code: code.OK,
		Message: "ok",
		Data: data,
	})
	c.Abort()
	return
}



func error_(c *gin.Context, status code.Code, err ...error) {
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(200,ApiResponse {
		Code: status,
		Message: status.Message(err...),
		Data: nil,
	})
	c.Abort()
	return
}

type ApiResponse struct {
	Code code.Code `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
