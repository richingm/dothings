package main

import (
	"github.com/gin-gonic/gin"
	"richingm/dothings/configs"
	"richingm/dothings/internal/repo"
	"richingm/dothings/middleware"
	"richingm/dothings/router"
)

func main() {
	config := configs.InitConfig()
	db := repo.InitDb(config)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.ConfigMiddleware(config))
	r.Use(middleware.DbMiddleware(db))
	router.InitRouter(r)
	r.Run(config.Server.HTTP.Addr)
}
