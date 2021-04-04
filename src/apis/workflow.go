package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jssyzhoulei/workflow/logger"
	"github.com/jssyzhoulei/workflow/src/models"
	"net/http"
)

func (a Apis) AddWorkflow(c *gin.Context) {
	var data = new(models.WorkFLow)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("GroupAdd 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}
	err = a.CreateFlow(data)
	if err != nil {
		log.Logger().Error("auth user error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", nil, false)
}

func (a Apis) ListWorkflow(c *gin.Context) {
	//userId := c.DefaultQuery("user_id", "0")

	res, err := a.ListFlow()
	if err != nil {
		log.Logger().Error("auth user error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", res, false)
}

func (a Apis) UpdateWorkflow(c *gin.Context) {
	var data = new(models.WorkFLow)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("GroupAdd 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}
	err = a.UpdateFlow(data)
	if err != nil {
		log.Logger().Error("auth user error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", nil, false)
}

func (a Apis) DelWorkflow(c *gin.Context) {
	var data = new(models.WorkFLow)

	err := c.BindJSON(data)
	if err != nil {
		log.Logger().Info(fmt.Sprintf("GroupAdd 参数解析错误: %s", err.Error()))
		response(c, http.StatusBadRequest, "参数解析错误", nil, false)
		return
	}
	err = a.DelFlow(data)
	if err != nil {
		log.Logger().Error("auth user error: " + err.Error())
		response(c, http.StatusBadRequest, "server error", nil, false)
		return
	}
	response(c, http.StatusOK, "success", nil, false)
}
