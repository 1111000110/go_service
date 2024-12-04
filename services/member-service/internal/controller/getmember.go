package controller

import (
	"context"
	"github.com/1111000110/go_service/services/member-service/internal/service"
	"github.com/1111000110/go_service/shared/proto/api/memberapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetmemberByMid(c *gin.Context) {
	param := memberapi.GetMemberByMidRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	data, err := service.GetMemberByMid(ctx, param.Mid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, data)
}
