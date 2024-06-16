package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-mongo/config"
	"github.com/prateek69/go-mongo/router.go"
)

func main() {
	config.LoadConfig()
	r := gin.New()
	router.UserRouter(r)
	r.Run(":" + config.AppConfig.Port)
}
