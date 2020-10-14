package apis

import (
	"context"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userApiInterface interface {
	AddUserApi(ctx *gin.Context)
	GetUserByIDApi(ctx *gin.Context)
}

type userApi struct {
	userService services.UserServiceInterface
}

func NewUserApi(userService services.UserServiceInterface) userApiInterface {
	return &userApi{
		userService: userService,
	}
}

func responses(ctx *gin.Context, status int, message string, data interface{}){
	if data == nil {
		data = ""
	}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.JSON(status, map[string]interface{} {
		"code" : status,
		"message" : message,
		"data" : data,
	})
	ctx.Abort()
	return
}

// AddUserApi 添加用户API
func (u *userApi) AddUserApi(ctx *gin.Context) {
	var data = new(pb_user_v1.UserProto)
	if err := ctx.BindJSON(data); err != nil{
		responses(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	_, err := u.userService.AddUserSvc(context.Background(), data)
	if err != nil{
		log.Logger().Error("创建用户失败:" + err.Error())
		responses(ctx, http.StatusBadRequest, "创建用户失败", nil)
		return
	}
	responses(ctx, http.StatusOK, "创建用户成功", nil)
}

func (u *userApi) GetUserByIDApi(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("uid"))

	if err != nil {
		responses(ctx, http.StatusBadRequest, "用户ID类型错误", nil)
		return
	}

	user, err := u.userService.GetUserByIDSvc(context.Background(), uid)

	if err != nil {
		log.Logger().Error("获取用户错误：" + err.Error())
		responses(ctx, http.StatusBadRequest, "获取用户失败", nil)
		return
	}

	responses(ctx, http.StatusOK, "获取用户成功", user)
}



