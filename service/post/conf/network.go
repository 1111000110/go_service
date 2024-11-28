package conf
import (
    "github.com/gin-gonic/gin"
    "log"
    "bytes"
    "io"
    "os"
    "path/filepath"
    "time"
)
func requestLogger() gin.HandlerFunc {
    date := time.Now().Format("2006-01-02")
    logDir := "logs"
    logFile := filepath.Join(logDir, date+".log")
    if _, err := os.Stat(logDir); os.IsNotExist(err) {
        if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
            log.Fatalf("Failed to create log directory: %v", err)
        }
    }
    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    logger := log.New(file, "", log.LstdFlags)
    return func(c *gin.Context) {
        body,_:=c.GetRawData()
        logger.Printf("IP: %s, Method: %s, Path: %s, Headers:%s, Request Body:", c.ClientIP(),c.Request.Method, c.Request.URL.Path,c.Request.Header,string(body))
        c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
        c.Next()
    }
}
func NetworkInit(r *gin.Engine){
	r.Use(requestLogger())
}