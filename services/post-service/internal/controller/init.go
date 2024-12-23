package controller

import (
	"github.com/gin-gonic/gin"
)

var postRoutes = map[string]gin.HandlerFunc{
	"/getpostsbypids":  GetPostsByPids,
	"/getpostsbymids":  GetPostsByMids,
	"/deletepostbypid": DeletePostByPid,
	"/deletepostbymid": DeletePostByMid,
	"/createpost":      CreatePost,
}
var testRoutes = map[string]gin.HandlerFunc{
	"/get": get,
}

func Init(r *gin.Engine) {
	apiGroup := r.Group("/post")
	for path, handler := range postRoutes {
		apiGroup.POST(path, handler)
	}
	apitest := r.Group("/test")
	for path, handler := range testRoutes {
		apitest.GET(path, handler)
	}
}
