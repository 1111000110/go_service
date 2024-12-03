package controller

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/service"
	"github.com/1111000110/go_service/shared/proto/api/postapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetPostsByPids(c *gin.Context) {
	param := &postapi.GetPostsByPidsRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	data, err := service.GetPostByPids(ctx, param.Pids) // 假设 param.Pids 是传递的多个 Post ID
	if err != nil || data == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := postapi.GetPostsByPidsResponse{
		Posts: data,
	}
	c.JSON(http.StatusOK, resp)
}
func GetPostsByMids(c *gin.Context) {
	param := &postapi.GetPostsByMidsRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	data, err := service.GetPostByMids(ctx, param.Mids)
	if err != nil || data == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := postapi.GetPostsByPidsResponse{
		Posts: data,
	}
	c.JSON(http.StatusOK, resp)
}
