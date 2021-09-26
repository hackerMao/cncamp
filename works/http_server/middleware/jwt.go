package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// Auth token验证中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 接收客户端request，并将request中带的header写入response header
		token := ctx.Request.Header.Get("Authorization")
		fmt.Printf("get token: %v\n", token)
		if token != "never to guess the secret!" {
			_, _ = ctx.Writer.WriteString("token invalid")
			ctx.Abort()
			return
		}
		fmt.Println("check pass")
		ctx.Writer.Header().Set("token", token)
		// 读取当前系统的环境变量中的VERSION配置，并写入response header
		version := os.Getenv("VERSION")
		if version == "" {
			version = "v1"
		}
		ctx.Writer.Header().Set("version", version)
		ctx.Next()
	}
}
