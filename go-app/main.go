package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/common"
	"github.com/prateek96/app/server"
)

func main() {
	realm := os.Getenv("realm")
	if realm == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	dependency := &common.Dependency{
		Router: gin.Default(),
	}

	server.StartServer(dependency)
}
