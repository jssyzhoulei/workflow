package apis

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userApiInterface interface {
	AddUserApi(ctx *gin.Context)
	GetUserByIDApi(ctx *gin.Context)
	UpdateUserByIDApi(ctx *gin.Context)
	DeleteUserByIDApi(ctx *gin.Context)
}

type userApi struct {
	userService services.UserServiceInterface
}

func NewUserApi(userService services.UserServiceInterface) userApiInterface {
	return &userApi{
		userService: userService,
	}
}

// AddUserApi 添加用户API
func (u *userApi) AddUserApi(ctx *gin.Context) {
	var (
		user models.User
	)
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("add user request param error : %s", err.Error()))
		response(ctx, http.StatusBadRequest, "param error", nil)
		return
	}
	_, err = u.userService.AddUserSvc(context.Background(), user)
	if err != nil {
		log.Logger().Error("add user error: " + err.Error())
		response(ctx, http.StatusBadRequest, "server error", nil)
		return
	}
	response(ctx, http.StatusOK, "success", nil)
	return
}

// GetUserByIDApi 获取用户详请
func (u *userApi) GetUserByIDApi(ctx *gin.Context){
	id := ctx.Param("id")
	ID,err := strconv.Atoi(id)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("get user request param error: %s", err.Error()))
		error_(ctx, 201, err)
		return
	}
	_, err = u.userService.GetUserByIDSvc(context.Background(), ID)
	if err != nil {
		log.Logger().Error("get user error: " + err.Error())
		error_(ctx, 201, err)
		return
	}
	success_(ctx, nil)
	return
}

// UpdateUserByIDApi 修改用户API
func (u *userApi) UpdateUserByIDApi(ctx *gin.Context) {

	var data = new(models.User)

	err := ctx.BindJSON(data)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("update user request param error : %s", err.Error()))
		response(ctx, http.StatusBadRequest, "param error", nil)
		return
	}
	_, err = u.userService.UpdateUserByIDSvc(context.Background(), *data)
	if err != nil {
		log.Logger().Error("update user error: " + err.Error())
		response(ctx, http.StatusBadRequest, "server error", nil)
		return
	}
	response(ctx, http.StatusOK, "success", nil)
	return
}

// DeleteUserByIDApi 删除用户API
func (u *userApi) DeleteUserByIDApi(ctx *gin.Context){
	var data = new(models.User)

	err := ctx.BindJSON(data)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("delete user request param error : %s", err.Error()))
		response(ctx, http.StatusBadRequest, "param error", nil)
		return
	}
	_, err = u.userService.DeleteUserByIDSvc(context.Background(), data.ID)
	if err != nil {
		log.Logger().Error("delete user error: " + err.Error())
		response(ctx, http.StatusBadRequest, "server error", nil)
		return
	}
	response(ctx, http.StatusOK, "success", nil)
	return
}


