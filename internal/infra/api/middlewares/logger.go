package middlewares

import (
	"time"

	"github.com/elaurentium/vaultui/internal/utils/log"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		latency := end.Sub(start)
		status := c.Writer.Status()

		log.Info("Request: method=%s path=%s status=%d latency=%s",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency.String(),
		)
	}
}
