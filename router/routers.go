package router

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {})
}
