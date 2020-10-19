package apis

import (
	"context"
	"encoding/base64"
	"fmt"
	"gitee.com/grandeep/org-svc/src/apis/code"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/tealeg/xlsx"
)

type userApiInterface interface {
	AddUserApi(ctx *gin.Context)
	GetUserByIDApi(ctx *gin.Context)
	UpdateUserByIDApi(ctx *gin.Context)
	DeleteUserByIDApi(ctx *gin.Context)
	ImportUser(ctx *gin.Context)
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
		response(ctx, http.StatusBadRequest, "param error", nil, false)
		return
	}
	_, err = u.userService.AddUserSvc(context.Background(), user)
	if err != nil {
		log.Logger().Error("add user error: " + err.Error())
		response(ctx, http.StatusBadRequest, "操作失败", nil, false)
		return
	}
	response(ctx, http.StatusOK, "操作成功", nil, false)
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
	data, err := u.userService.GetUserByIDSvc(context.Background(), ID)
	if err != nil {
		log.Logger().Error("get user error: " + err.Error())
		error_(ctx, 201, err)
		return
	}
	success_(ctx, data)
	return
}

// UpdateUserByIDApi 修改用户API
func (u *userApi) UpdateUserByIDApi(ctx *gin.Context) {

	var data = new(models.User)

	err := ctx.BindJSON(data)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("update user request param error : %s", err.Error()))
		response(ctx, http.StatusBadRequest, "param error", nil, false)
		return
	}
	_, err = u.userService.UpdateUserByIDSvc(context.Background(), *data)
	if err != nil {
		log.Logger().Error("update user error: " + err.Error())
		response(ctx, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(ctx, http.StatusOK, "success", nil, false)
	return
}

// DeleteUserByIDApi 删除用户API
func (u *userApi) DeleteUserByIDApi(ctx *gin.Context){
	var data = new(models.User)

	err := ctx.BindJSON(data)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("delete user request param error : %s", err.Error()))
		response(ctx, http.StatusBadRequest, "param error", nil, false)
		return
	}
	_, err = u.userService.DeleteUserByIDSvc(context.Background(), data.ID)
	if err != nil {
		log.Logger().Error("delete user error: " + err.Error())
		response(ctx, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(ctx, http.StatusOK, "success", nil, false)
	return
}


//导入用户
func (u *userApi) ImportUser(ctx *gin.Context) {
	var (
		importUserRequest models.ImportUserRequest
		content []byte
		file *xlsx.File
		users = &pb_user_v1.AddUsersRequest{}
	)
	err := ctx.BindJSON(&importUserRequest)
	if err != nil {
		error_(ctx, code.PARAMS_ERROR)
		return
	}
	content, err = base64.StdEncoding.DecodeString(importUserRequest.Content)
	if err != nil {
		error_(ctx, code.PARAMS_ERROR)
		return
	}
	file, err = xlsx.OpenBinary(content)
	if err != nil {
		error_(ctx, code.PARAMS_ERROR)
		return
	}
	for _, sheet := range file.Sheets {
		//验证模板是否正确
		for k, row := range sheet.Rows {
			if k == 0 {
				if len(row.Cells) >= 4 {
					if row.Cells[0].Value != "*用户名" || row.Cells[1].Value != "*登录名" || row.Cells[2].Value != "*密码" || row.Cells[3].Value != "手机号" {
						error_(ctx, code.XlsxError)
						return
					}
				}
				continue
			}
			if row.Cells[0].Value != "" && row.Cells[1].Value != "" && row.Cells[2].Value != "" {
				var (
					user pb_user_v1.UserProto
				)
				user.UserName = row.Cells[0].Value
				user.Password = row.Cells[2].Value
				user.LoginName = row.Cells[1].Value
				user.Mobile,_ = strconv.ParseInt(row.Cells[1].Value, 10, 64)
				user.GroupId  = importUserRequest.GroupID
				for _, v := range importUserRequest.RoleID {
					user.RoleIds = append(user.RoleIds, &pb_user_v1.Index{Id: v})
				}
				users.Users = append(users.Users, &user)
			}
		}
	}
	users.IsCover = importUserRequest.IsCover
	_ , err = u.userService.AddUsersSvc(context.Background(), users)
	fmt.Println(err)
	if err != nil {
		error_(ctx, code.SVC_ERROR)
		return
	}
	success_(ctx, nil)
}


