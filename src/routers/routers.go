package routers

import (
	"gitee.com/grandeep/org-svc/client"
	"gitee.com/grandeep/org-svc/src/apis"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	once sync.Once
	engine *gin.Engine
)

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		cost := time.Since(start)
			logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.RequestURI),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("cost", cost),
		)
	}
}

func Gin() *gin.Engine {
	once.Do(func() {
		engine = gin.New()
		engine.Use(GinLogger(log.Logger()))
	})
	return engine
}

func Routers(e *gin.Engine) {
	o := client.NewOrgServiceClient([]string{"172.18.28.226:2379"}, 3, time.Second)
	api := apis.NewApis(o)
	g := e.Group("/apis/v1")
	userApiRouters(g, api)
	permissionApiRouters(g, api)
	groupAPIRouters(g, api)
}

