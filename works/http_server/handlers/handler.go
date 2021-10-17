package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheckHandler 健康监测处理器
func HealthCheckHandler(c *gin.Context) {
	c.Writer.WriteString("health")
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "<h2>hello gopher!</h2>",
	})
}
