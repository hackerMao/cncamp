package handlers

import (
	"github.com/gin-gonic/gin"
)

// HealthCheckHandler 健康监测处理器
func HealthCheckHandler(c *gin.Context) {
	c.Writer.WriteString("health")
}
