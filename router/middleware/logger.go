package middleware

import (
	"gin-example/application/library"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		library.Logger.Info("接收到请求：",
			zap.String("path", c.FullPath()),
			zap.String("referer", c.Request.Referer()))

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		library.Logger.Info("请求结束",
			zap.Duration("consume", latency),
			zap.Int("status", status))
	}
}
