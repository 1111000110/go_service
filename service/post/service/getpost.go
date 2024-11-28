package service

import (
	"github.com/1111000110/go_service/api/api/postapi"
	"github.com/1111000110/go_service/api/apidata/postapidata"
	"github.com/gin-gonic/gin"
)

func GetPost(ctx *gin.Context, param *postapi.GetPostParam) (*postapidata.Post, error) {
	data := &postapidata.Post{
		Pid:  1,
		Mid:  2,
		Tid:  3,
		Ct:   4,
		Ut:   5,
		Text: "test",
		Img:  "6",
	}
	return data, nil
}
