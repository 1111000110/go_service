package main

import (
    "github.com/1111000110/go_service/service/post/controller"
    "log"
    "net/http"
	"fmt"
)
func main(){
	controller:=&controller.Controller{}
	controller.InitRoutes();//加载路由
	fmt.Println("Hello")
    // 定义服务器要监听的端口
    port := "8080"
    // 启动服务器
    log.Printf("Server starting on port %s\n", port)
    if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}