package main

import (
	"github.com/1111000110/go_service/service/post/conf"
	"github.com/1111000110/go_service/service/post/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	conf.NetworkInit(r) //加载配置
	controller.Init(r)  //加载路由
	r.Run(":8080")
}
