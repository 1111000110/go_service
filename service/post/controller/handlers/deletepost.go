package handlers

import (
	"github.com/1111000110/go_service/api/api/postapi"
	"github.com/1111000110/go_service/service/post/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeletePost(c *gin.Context) {
	param := &postapi.DeletePostParam{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data, err := service.DeletePost(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	resp := postapi.DeletePostResp{
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}
