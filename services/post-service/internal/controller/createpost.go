package controller

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/service"
	"github.com/1111000110/go_service/shared/models"
	"github.com/1111000110/go_service/shared/proto/api/postapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {
	param := &postapi.CreatePostRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 设置超时时间
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	data, err := service.CreatePost(ctx, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	resp := postapi.CreatePostResponse{
		Post: models.Post{
			Tid:   data.Tid,
			Mid:   data.Mid,
			Pid:   data.Pid,
			Title: data.Title,
			Desc:  data.Desc,
			Ct:    data.Ct,
			Ut:    data.Ut,
		},
	}
	c.JSON(http.StatusOK, resp)
}
