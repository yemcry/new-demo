package router

import (
	"net/http"
	"server/ws"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	// 设置 CORS 中间件
	r.Use(CORSMiddleware())
	r.GET("/ws", ws.HandleWebSocket)
	UserLogin(r)

	r.Run(":12345")
}

// CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")                            // 允许的来源
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")          // 允许的方法
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization") // 允许的请求头

		// 处理预检请求
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // 返回204 No Content
			return
		}

		c.Next() // 继续处理请求
	}
}
