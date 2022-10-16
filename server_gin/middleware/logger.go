package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var ginLogger = gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.TimeStamp.Format(time.RFC3339),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
})

func GetGinLogger() gin.HandlerFunc {
	return ginLogger
}
