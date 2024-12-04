package controller

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/service"
	"github.com/1111000110/go_service/shared/proto/api/postapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func DeletePostByMid(c *gin.Context) {
	param := &postapi.DeletePostByMidRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 设置超时时间
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	err := service.DeletePostByMid(ctx, param.Mid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := postapi.DeletePostByMidResponse{
		IsDeleted: true,
	}
	c.JSON(http.StatusOK, resp)
}
func DeletePostByPid(c *gin.Context) {
	param := &postapi.DeletePostByPidRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 设置超时时间
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	err := service.DeletePostByPid(ctx, param.Mid, param.Pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := postapi.DeletePostByPidResponse{
		IsDeleted: true,
	}
	c.JSON(http.StatusOK, resp)
}
