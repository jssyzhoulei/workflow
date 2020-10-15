package apis

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/gin-gonic/gin"
)

type permissionApiInterface interface {
	AddPermissionApi(ctx *gin.Context)
}

type permissionApi struct {
	permissionService services.PermissionServiceInterface
}

func NewPermissionApi(permissionService services.PermissionServiceInterface) permissionApiInterface {
	return &permissionApi{
		permissionService:  permissionService,
	}
}

func (p *permissionApi) AddPermissionApi(ctx *gin.Context) {
	var (
		permission models.Permission
	)
	err := ctx.BindJSON(&permission)
	if err != nil {
		error_(ctx, 201, err)
		return
	}
	_, err = p.permissionService.AddPermissionSvc(context.Background(), permission)
	if err != nil {
		error_(ctx, 201, err)
		return
	}
	success_(ctx, nil)
	return
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