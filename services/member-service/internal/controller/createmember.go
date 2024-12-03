package controller

import (
	"context"
	"github.com/1111000110/go_service/services/member-service/internal/service"
	"github.com/1111000110/go_service/shared/proto/api/memberapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateMember(c *gin.Context) {
	param := &memberapi.CreateMemberRequest{}
	if err := c.ShouldBindJSON(param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()
	resp, err := service.CreateMember(ctx, param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp)
}
