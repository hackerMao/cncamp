package routers

import (
	"github.com/gin-gonic/gin"
	"http_server/handlers"
	"http_server/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode("debug")
	engine := gin.New()
	// Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
	engine.Use(gin.Logger(), gin.Recovery())

	group := engine.Group("/")
	{
		// 健康检测
		group.GET("/healthz", handlers.HealthCheckHandler)
	}

	api := engine.Group("/api", middleware.Auth())
	{
		api.GET("/index", handlers.Index)
	}

	return engine
}
