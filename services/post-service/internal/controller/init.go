package controller

import (
	"github.com/gin-gonic/gin"
)

var postRoutes = map[string]gin.HandlerFunc{
	"/getpostsbypids":  GetPostsByPids,
	"/getpostsbymids":  GetPostsByMids,
	"/deletepostbypid": DeletePostByPid,
	"/createpost":      CreatePost,
	"/":                test,
}

func Init(r *gin.Engine) {
	apiGroup := r.Group("/post")
	for path, handler := range postRoutes {
		apiGroup.POST(path, handler)
	}
	apiGroup.GET("/", test)
}
