package main

import (
    "github.com/1111000110/go_service/service/post/conf"
    "github.com/gin-gonic/gin"
    "github.com/1111000110/go_service/service/post/controller"
)
type resp struct{
    Str string `json:"str"`
}

func main()  {
    r := gin.Default()
    conf.NetworkInit(r)//加载配置
    controller.Init(r)//加载路由
    r.Run(":8080")
}