package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"richingm/dothings/configs"
)

const ConfigContextKey string = "config"

func ConfigMiddleware(config *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ConfigContextKey, config)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
