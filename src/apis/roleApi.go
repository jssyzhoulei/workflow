package apis

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/logger"
	"gitee.com/grandeep/org-svc/src/apis/code"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type RoleApiInterface interface {
	AddRoleApi(ctx *gin.Context)
	UpdateRoleApi(ctx *gin.Context)
	DeleteRoleApi(ctx *gin.Context)
	QueryRoleApi(ctx *gin.Context)
	QueryRolesApi(ctx *gin.Context)
	GetMenuCascade(c *gin.Context)
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
		log.Logger.Warn(fmt.Sprintf("add role request param error with data : %+v ", data))
		response(c, http.StatusBadRequest, "param error", nil, false)
		return
	}
	_, err = r.roleService.AddRoleSvc(context.Background(), data)
	if err != nil {
		log.Logger.Error("create role error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", nil, false)
}

func (r *roleApi) UpdateRoleApi(c *gin.Context) {

	var data = new(models.CreateMenuPermRequest)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger.Warn(fmt.Sprintf("update role request param error : %s", err.Error()))
		response(c, http.StatusBadRequest, "param error", nil, false)
		return
	}
	_, err = r.roleService.UpdateRoleSvc(context.Background(), data)
	if err != nil {
		log.Logger.Error("update role error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", nil, false)
}

func (r *roleApi) DeleteRoleApi(c *gin.Context) {

	var idStr = c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response(c, http.StatusBadRequest, "param error", nil, false)
		return
	}
	_, err = r.roleService.DeleteRoleSvc(context.Background(), id)
	if err != nil {
		if strings.Contains(err.Error(), "relation user") {
			response(c, 201, "请先解除绑定用户角色", "", false)
			return
		}
		log.Logger.Error("delete role error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", nil, false)
}

func (r *roleApi) QueryRoleApi(c *gin.Context) {

	var idStr = c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id == 0 {
		response(c, http.StatusBadRequest, "param error", nil, false)
		return
	}
	resp, err := r.roleService.QueryRoleSvc(context.Background(), id)
	if err != nil || len(resp.MenuPerms) == 0 {
		log.Logger.Error("query role error ")
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	resp.ID = resp.MenuPerms[0].RoleID
	response(c, http.StatusOK, "success", resp, false)
}

func (r *roleApi) QueryRolesApi(c *gin.Context) {

	var data = new(pb_user_v1.RolePageRequestProto)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger.Warn(fmt.Sprintf("query roles request param error : %s", err.Error()))
		response(c, http.StatusBadRequest, "param error", nil, false)
		return
	}
	resp, err := r.roleService.QueryRolesSvc(context.Background(), data)
	if err != nil {
		log.Logger.Error("query roles error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", resp, false)
}

func (r *roleApi) GetMenuCascade(c *gin.Context) {
	module := c.Query("module")
	m, err := strconv.Atoi(module)
	if err != nil {
		error_(c, code.PARAMS_ERROR)
		return
	}
	cascades, err := r.roleService.MenuTreeSvc(context.Background(), models.MenuModule(m))
	if err != nil {
		error_(c, code.SVC_ERROR)
		return
	}
	success_(c, cascades.Cascades)
	return
}
