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

func GetPostsByPids(c *gin.Context) {
	// 绑定请求参数
	param := &postapi.GetPostsByPidsRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		// 参数绑定失败，返回 400 错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 设置超时时间
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	// 调用 service 获取数据
	data, err := service.GetPostByPids(ctx, param.Pids) // 假设 param.Pids 是传递的多个 Post ID
	if err != nil || data == nil {
		// 如果查询失败或者数据为空，返回 500 错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 构造响应数据
	var posts []models.Post
	for _, item := range *data {
		posts = append(posts, models.Post{
			Pid:   item.Pid,
			Title: item.Title,
			Desc:  item.Desc,
			Mid:   item.Mid,
			Tid:   item.Tid,
			Ct:    item.Ct,
			Ut:    item.Ut,
		})
	}

	// 返回查询结果，包含多个 Post 对象
	resp := postapi.GetPostsByPidsResponse{
		Posts: posts, // 返回多个 Post
	}
	// 返回 200 OK 响应
	c.JSON(http.StatusOK, resp)
}
func GetPostsByMids(c *gin.Context) {
	// 绑定请求参数
	param := &postapi.GetPostsByMidsRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		// 参数绑定失败，返回 400 错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 设置超时时间
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	// 调用 service 获取数据
	data, err := service.GetPostByMids(ctx, param.Mids) // 假设 param.Pids 是传递的多个 Post ID
	if err != nil || data == nil {
		// 如果查询失败或者数据为空，返回 500 错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 构造响应数据
	var posts []models.Post
	for _, item := range *data {
		posts = append(posts, models.Post{
			Pid:   item.Pid,
			Title: item.Title,
			Desc:  item.Desc,
			Mid:   item.Mid,
			Tid:   item.Tid,
			Ct:    item.Ct,
			Ut:    item.Ut,
		})
	}

	// 返回查询结果，包含多个 Post 对象
	resp := postapi.GetPostsByPidsResponse{
		Posts: posts, // 返回多个 Post
	}
	// 返回 200 OK 响应
	c.JSON(http.StatusOK, resp)
}
