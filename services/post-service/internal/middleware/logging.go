package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"time"
)

// LoggingMiddleware 用于记录请求的日志（IP、请求头、请求体等）
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 获取客户端的 IP 地址
		clientIP := c.ClientIP()

		// 获取请求方法和 URL
		method := c.Request.Method
		url := c.Request.URL.Path

		// 获取请求头
		headers := c.Request.Header

		// 读取请求体
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			c.JSON(500, gin.H{"error": "Unable to read request body"})
			c.Abort()
			return
		}

		// 保存原始请求体，供后续使用
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 打印日志（记录请求的 IP、头部、请求体）
		log.Printf("IP: %s, Method: %s, URL: %s, Headers: %v, Body: %s", clientIP, method, url, headers, string(bodyBytes))

		// 在请求前执行，继续处理请求
		c.Next()

		// 计算请求处理耗时
		duration := time.Since(start)

		// 获取响应的状态码
		statusCode := c.Writer.Status()

		// 打印处理后的日志信息（请求耗时、响应状态码等）
		log.Printf("Method: %s, URL: %s, Status: %d, Duration: %v", method, url, statusCode, duration)
	}
}
