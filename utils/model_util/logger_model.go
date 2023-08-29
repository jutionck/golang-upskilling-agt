package modelutil

import "github.com/gin-gonic/gin"

type RequestLog struct {
	Method     string
	StatusCode int
	ClientIP   string
	Path       string
	UserAgent  string
}

type responsWriter struct {
	gin.ResponseWriter
	status int
	body   []byte
}

type ResponseLog struct {
	StatusCode   int
	ResponseBody string
}

func NewResponseLog(rw gin.ResponseWriter) *responsWriter {
	return &responsWriter{ResponseWriter: rw}
}

func (r *responsWriter) Write(data []byte) (int, error) {
	r.body = append(r.body, data...)
	return r.ResponseWriter.Write(data)
}

func (r *responsWriter) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responsWriter) Body() string {
	return string(r.body)
}

func (r *responsWriter) Status() int {
	return r.status
}
