package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const DbContextKey string = "db"

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), DbContextKey, db)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
