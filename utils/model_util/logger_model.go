package modelutil

import "github.com/gin-gonic/gin"

type RequestLog struct {
	Method     string
	StatusCode int
	ClientIP   string
	Path       string
	UserAgent  string
}

type ResponseLog struct {
	gin.ResponseWriter
	StatusCode   int
	ResponseBody []byte
}
