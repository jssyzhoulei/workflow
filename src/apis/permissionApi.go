package apis

import (
	"context"
	"gitee.com/grandeep/org-svc/src/apis/code"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type permissionApiInterface interface {
	AddPermissionApi(ctx *gin.Context)
	AddMenuApi(c *gin.Context)
	GetMenuCascade(c *gin.Context)
	GetPermissionByID(c *gin.Context)
	DeletePermissionByID(c *gin.Context)
	UpdatePermission(c *gin.Context)
}

type permissionApi struct {
	permissionService services.PermissionServiceInterface
}

func NewPermissionApi(permissionService services.PermissionServiceInterface) permissionApiInterface {
	return &permissionApi{
		permissionService: permissionService,
	}
}

func (p *permissionApi) AddPermissionApi(ctx *gin.Context) {
	var (
		permission models.Permission
	)
	err := ctx.BindJSON(&permission)
	if err != nil {
		error_(ctx, code.PARAMS_ERROR)
		return
	}
	_, err = p.permissionService.AddPermissionSvc(context.Background(), permission)
	if err != nil {
		error_(ctx, code.SVC_ERROR)
		return
	}
	success_(ctx, nil)
	return
}

func (p *permissionApi) AddMenuApi(c *gin.Context) {
	var (
		menu models.Menu
	)
	err := c.BindJSON(&menu)
	if err != nil {
		error_(c, code.PARAMS_ERROR)
		return
	}
	_, err = p.permissionService.AddMenuSvc(context.Background(), menu)
	if err != nil {
		error_(c, code.SVC_ERROR)
		return
	}
	success_(c, nil)
	return
}

func (p *permissionApi) GetMenuCascade(c *gin.Context) {
	module := c.Query("module")
	m, err := strconv.Atoi(module)
	if err != nil {
		error_(c, code.PARAMS_ERROR)
		return
	}
	cascades, err := p.permissionService.GetMenuCascadeByModuleSvc(context.Background(), models.MenuModule(m))
	if err != nil {
		error_(c, code.SVC_ERROR)
		return
	}
	success_(c, cascades.Cascades)
	return
}

func (p *permissionApi) GetPermissionByID(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		error_(c, code.PARAMS_ERROR)
		return
	}
	permission, err := p.permissionService.GetPermissionByIDSvc(context.Background(), ID)
	if err != nil {
		error_(c, code.SVC_ERROR)
		return
	}
	success_(c, permission)
}

func (p *permissionApi) DeletePermissionByID(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		error_(c, code.PARAMS_ERROR)
		return
	}
	permission, err := p.permissionService.DeletePermissionByIDSvc(context.Background(), ID)
	if err != nil {
		error_(c, code.SVC_ERROR)
		return
	}
	success_(c, permission)
}

func (p *permissionApi) UpdatePermission(c *gin.Context) {
	var (
		permission models.Permission
	)
	err := c.BindJSON(&permission)
	if err != nil {
		error_(c, code.PARAMS_ERROR)
		return
	}
	_, err = p.permissionService.UpdatePermissionByIDSvc(context.Background(), permission)
	if err != nil {
		error_(c, code.SVC_ERROR)
		return
	}
	success_(c, nil)
	return
}
