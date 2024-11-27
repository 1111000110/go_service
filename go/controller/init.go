package controller
import(
    "net/http"
    "fmt"
)
type Controller struct{

}
// InitRoutes 初始化路由
func (c *Controller) InitRoutes() {
    http.HandleFunc("/",c.Handler)
	http.HandleFunc("/post", c.postHandler)
	http.HandleFunc("/review", c.reviewHandler)
	http.HandleFunc("/topic", c.topicHandler)
	http.HandleFunc("/member", c.memberHandler)
}
func (c *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "刘雪倩是个大笨蛋！！！！")
    fmt.Println(r.FormValue)
}
// postHandler 处理 POST 路由请求
func (c *Controller) postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling POST request")
    fmt.Println(r.FormValue)
}

// reviewHandler 处理 REVIEW 路由请求
func (c *Controller) reviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling REVIEW request")
    fmt.Println(r.FormValue)
}

// topicHandler 处理 TOPIC 路由请求
func (c *Controller) topicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling TOPIC request")
    fmt.Println(r.FormValue)
}

// memberHandler 处理 MEMBER 路由请求
func (c *Controller) memberHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling MEMBER request")
    fmt.Println(r.FormValue)
}