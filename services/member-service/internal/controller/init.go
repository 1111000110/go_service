package controller

import (
	"github.com/gin-gonic/gin"
)

var postRoutes = map[string]gin.HandlerFunc{
	"/createmember":   CreateMember,
	"/deletemember":   DeleteMember,
	"/getmemberbymid": GetmemberByMid,
}

func Init(r *gin.Engine) {
	apiGroup := r.Group("/member")
	for path, handler := range postRoutes {
		apiGroup.POST(path, handler)
	}
}
