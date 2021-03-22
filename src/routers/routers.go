package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jssyzhoulei/workflow/logger"
	"github.com/jssyzhoulei/workflow/src/apis"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	once      sync.Once
	engine    *gin.Engine
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, Language, From")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Gin() *gin.Engine {
	once.Do(func() {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(Cors())
		engine.Use(GinLogger(log.Logger()))
	})
	return engine
}

func Routers(e *gin.Engine, api *apis.Apis) {
	g := e.Group("/apis/v1/")
	workApiRouters(g, api)
}
