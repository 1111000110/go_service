package main

import (
	"fmt"
	"github.com/1111000110/go_service/services/member-service/internal/config"
	"github.com/1111000110/go_service/services/member-service/internal/controller"
	"github.com/1111000110/go_service/services/member-service/internal/middleware"
	"github.com/1111000110/go_service/services/member-service/internal/repository/mongo"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	mongo.Init() //初始化mongo
	// 设置 Gin 路由
	router := gin.Default()
	// 中间日志
	router.Use(middleware.LoggingMiddleware())
	// 注册路由
	controller.Init(router)
	// 启动服务
	serverPort := config.AppConfig.Server.Port
	if err := router.Run(fmt.Sprintf(":%d", serverPort)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
