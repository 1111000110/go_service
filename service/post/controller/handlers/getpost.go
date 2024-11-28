package handlers

import (
	"github.com/1111000110/go_service/api/api/postapi"
	"github.com/1111000110/go_service/service/post/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPost(ctx *gin.Context) {
	param := &postapi.GetPostParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data, err := service.GetPost(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	resp := postapi.GetPostResp{
		Data: data,
	}
	ctx.JSON(http.StatusOK, resp)
}
