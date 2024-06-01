package middlewares

import (
	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// statusCodeColor := params.StatusCodeColor()
		// methodColor := params.MethodColor()
		// resetColor := params.ResetColor()

		// return fmt.Sprintf("%s\t%s%s%s\t[%s]\t%s\t%s|%s %d %s|\t%s\t%s\n",
		// 	params.TimeStamp.Format("2006-01-02 15:04:05.000000"),
		// 	methodColor,
		// 	params.Method,
		// 	resetColor,
		// 	params.ClientIP,
		// 	params.Path,
		// 	params.Request.Proto,
		// 	statusCodeColor,
		// 	params.StatusCodeColor(),
		// 	resetColor,
		// 	params.Latency,
		// 	params.ErrorMessage,
		// )
		return ""
	})
}
