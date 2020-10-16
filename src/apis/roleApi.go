package apis

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleApiInterface interface {
	AddRoleApi(ctx *gin.Context)
	UpdateRoleApi(ctx *gin.Context)
	DeleteRoleApi(ctx *gin.Context)
}

type roleApi struct {
	roleService services.RoleServiceI
}

func NewRoleApi(roleService services.RoleServiceI) RoleApiInterface {
	return &roleApi{
		roleService: roleService,
	}
}

func (r *roleApi) AddRoleApi(c *gin.Context) {

	var data = new(models.CreateMenuPermRequest)

	err := c.BindJSON(data)
	if err != nil || !data.Check() {
		log.Logger().Warn(fmt.Sprintf("add role request param error with data : %+v ", data))
		response(c, http.StatusBadRequest, "param error", nil)
		return
	}
	_, err = r.roleService.AddRoleSvc(context.Background(), *data)
	if err != nil {
		log.Logger().Error("create role error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil)
		return
	}
	response(c, http.StatusOK, "success", nil)
}

func (r *roleApi) UpdateRoleApi(c *gin.Context) {

	var data = new(models.CreateMenuPermRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Warn(fmt.Sprintf("update role request param error : %s", err.Error()))
		response(c, http.StatusBadRequest, "param error", nil)
		return
	}
	_, err = r.roleService.UpdateRoleSvc(context.Background(), *data)
	if err != nil {
		log.Logger().Error("update role error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil)
		return
	}
	response(c, http.StatusOK, "success", nil)
}

func (r *roleApi) DeleteRoleApi(c *gin.Context) {

	var data = new(models.CreateMenuPermRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Warn(fmt.Sprintf("delete role request param error : %s", err.Error()))
		response(c, http.StatusBadRequest, "param error", nil)
		return
	}
	_, err = r.roleService.DeleteRoleSvc(context.Background(), *data)
	if err != nil {
		log.Logger().Error("delete role error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil)
		return
	}
	response(c, http.StatusOK, "success", nil)
}
