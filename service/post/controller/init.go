package controller
import (
    "github.com/gin-gonic/gin"
	"github.com/1111000110/go_service/service/post/controller/handler"
)
var postRoutes = map[string]gin.HandlerFunc{
	"/getpost": handler.getpost,
	"/deletepost": handler.deletepost,
	"/addpost": handler.addpost,
}
func Init(r *gin.Engine){
	apiGroup := r.Group("/post")
	for path, handler := range postRoutes {
        apiGroup.POST(path, handler)
    }
}