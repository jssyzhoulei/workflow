package apis

import (
	"gitee.com/grandeep/org-svc/client"
	"gitee.com/grandeep/org-svc/src/apis/code"
	"github.com/gin-gonic/gin"
)

type IApis interface {
	GetUserApis() userApiI
	GetPermission() permissionApiInterface
	GetRoleApis() RoleApiInterface
	GetGroupApis() groupApiInterface
}

type apis struct {
	userApiI
	permissionApiInterface permissionApiInterface
	RoleApiInterface
	groupApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		RoleApiInterface:NewRoleApi(o.GetRoleService()),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
		groupApiInterface: NewGroupApi(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
}

func (a *apis) GetRoleApis() RoleApiInterface {
	return a.RoleApiInterface
}

func (a *apis) GetGroupApis() groupApiInterface {
	return a.groupApiInterface
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
	Data interface{}
}