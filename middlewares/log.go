package middlewares

import (
	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start := time.Now()
		// path := c.Request.URL.Path

		c.Next()

		// cost := time.Since(start)
		// logger.Info(
		// 	1,
		// 	"%d %s:%s Cost:%v",
		// 	c.Writer.Status(),
		// 	c.Request.Method,
		// 	path,
		// 	cost,
		// )
	}
}
