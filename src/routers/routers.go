package routers

import (
	"gitee.com/grandeep/org-svc/client"
	"sync"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	once sync.Once
	engine *gin.Engine
)

func Gin() *gin.Engine {
	once.Do(func() {
		engine = gin.New()
	})
	return engine
}

func Routers() {
	o := client.NewOrgServiceClient([]string{"127.0.0.1:2379"}, 3, time.Second)
	g := Gin().Group("/apis/v1")
	userApiRouters(g, o)
}
