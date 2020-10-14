package apis

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/gin-gonic/gin"
)

type userApiI interface {
	AddUserApi(ctx *gin.Context)
}

type userApi struct {
	userService services.UserServiceI
}

func NewUserApi(userService services.UserServiceI) userApiI {
	return &userApi{
		userService: userService,
	}
}

func (u *userApi) AddUserApi(ctx *gin.Context) {
	fmt.Println(u.userService.AddUser(context.Background(), models.User{}))
}



