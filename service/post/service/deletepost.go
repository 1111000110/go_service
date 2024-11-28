package service

import (
	"fmt"
	"github.com/1111000110/go_service/api/api/postapi"
	"github.com/gin-gonic/gin"
)

func DeletePost(ctx *gin.Context, param *postapi.DeletePostParam) (bool, error) {
	pid := param.Pid
	fmt.Println("pid=", pid)
	return true, nil
}
