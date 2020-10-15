package apis

import (
	"gitee.com/grandeep/org-svc/client"
	"github.com/gin-gonic/gin"
)

type IApis interface {
	GetUserApis() userApiI
	GetPermission() permissionApiInterface
	GetGroupApis() groupApiInterface

}

type apis struct {
	userApiI
	permissionApiInterface permissionApiInterface
	groupApiInterface
}

func NewApis(o *client.OrgServiceClient) IApis {
	return &apis{
		userApiI: NewUserApi(o.GetUserService()),
		//groupApiI,NewGroupApi(o.),
		permissionApiInterface: NewPermissionApi(o.GetPermissionService()),
		groupApiInterface: NewGroupApi(o.GetGroupService()),
	}
}

func (a *apis) GetUserApis() userApiI {
	return a.userApiI
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
	c.JSON(200, map[string]interface{} {
		"code": 200,
		"message": "",
		"data": data,
	})
	c.Abort()
	return
}



func error_(c *gin.Context, status int, err error) {
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(200, map[string]interface{} {
		"code": status,
		"message": err.Error(),
		"data": nil,
	})
	c.Abort()
	return
}
