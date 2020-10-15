package apis

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/gin-gonic/gin"
)

type userApiInterface interface {
	AddUserApi(ctx *gin.Context)
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
		errord_(ctx, 201, err)
		return
	}
	_, err = u.userService.AddUserSvc(context.Background(), user)
	if err != nil {
		error_(ctx, 201, err)
		return
	}
	successd_(ctx, nil)
	return
}

func successd_(c *gin.Context, data interface{}) {
	if data == nil {
		data = ""
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(200, map[string]interface{} {
		"code" : 200,
		"message" : "",
		"data" : data,
	})
	c.Abort()
	return
}

func errord_(c *gin.Context, status int, err error) {
	c.Request.Header.Set("Content-Type", "application/json")
	c.JSON(200, map[string]interface{} {
		"code" : status,
		"message" : err.Error(),
		"data" : nil,
	})
	c.Abort()
	return
}



