package middleware

import (
	"github.com/gin-gonic/gin"
	loggerutil "github.com/jutionck/golang-upskilling-agt/utils/logger_util"
	modelutil "github.com/jutionck/golang-upskilling-agt/utils/model_util"
)

type LogMiddleware interface {
	Logger() gin.HandlerFunc
}

type logMiddleware struct {
	loggerService loggerutil.LoggerUtil
}

// LogRequest implements LogMiddleware.
func (l *logMiddleware) Logger() gin.HandlerFunc {
	err := l.loggerService.InitialLoggerFile()
	if err != nil {
		panic(err)
	}

	return func(ctx *gin.Context) {
		requestLog := modelutil.RequestLog{
			Method:     ctx.Request.Method,
			StatusCode: ctx.Writer.Status(),
			ClientIP:   ctx.ClientIP(),
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		}

		responseLog := modelutil.ResponseLog{
			ResponseWriter: ctx.Writer,
			StatusCode:     ctx.Writer.Status(),
		}

		// Untuk proses request seterusnya
		ctx.Writer = responseLog.ResponseWriter
		ctx.Next()

		switch {
		case ctx.Writer.Status() >= 400:
			l.loggerService.ReqLogError(requestLog)
			l.loggerService.ResLogError(responseLog)
		default:
			l.loggerService.ReqLogInfo(requestLog)
			l.loggerService.ResLogInfo(responseLog)
		}
	}
}

func NewLogMiddleware(loggerService loggerutil.LoggerUtil) LogMiddleware {
	return &logMiddleware{loggerService: loggerService}
}
