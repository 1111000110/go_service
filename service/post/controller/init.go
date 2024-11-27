package controller
import(
	"github.com/1111000110/go_service/api/postapi"
	"net/http"
	"fmt"
	"io"
    "encoding/json" // 导入 json 包
)
type Controller struct {

}
func (c *Controller) InitRoutes(){
	http.HandleFunc("/post",c.GetPostHandler)
}
func (c*Controller)GetPostHandler(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var jsonData postapi.PostParam
	if err := json.Unmarshal(body, &jsonData); err != nil {
		fmt.Fprintf(w, "刘雪倩是个大笨蛋！！！！")
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(jsonData)
	fmt.Println(jsonData.Str)
	  // 使用标准库 encoding/json 将结构体编码为 JSON
	  response, err := json.Marshal(jsonData)
	  if err != nil {
		  http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		  return
	  }
	  // 设置响应头 Content-Type 为 application/json
	  w.Header().Set("Content-Type", "application/json")
	  // 写入 JSON 响应体
	   // 设置状态码为 200 OK
	   w.WriteHeader(http.StatusOK)
	  w.Write(response)
}