package controller

import (
	"github.com/1111000110/go_service/service/post/controller/handlers"
	"github.com/gin-gonic/gin"
)

var postRoutes = map[string]gin.HandlerFunc{
	"/getpost":    handlers.GetPost,
	"/deletepost": handlers.DeletePost,
	"/addpost":    handlers.AddPost,
}

func Init(r *gin.Engine) {
	apiGroup := r.Group("/post")
	for path, handler := range postRoutes {
		apiGroup.POST(path, handler)
	}
}
