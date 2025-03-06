package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/common"
	"github.com/prateek96/app/controller"
	"github.com/prateek96/app/middleware"
)

func initializeRoutes(dependency *common.Dependency) {
	dependency.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	api := dependency.Router.Group("/api")
	{
		postApi := api.Group("v1/path1", middleware.PostMiddleware(dependency))
		postApi.POST("/post", controller.PostController(dependency))

		getApi := api.Group("/v1/path2", middleware.GetMiddleware(dependency))
		getApi.GET("/get", controller.GetController(dependency))
	}
}

func StartServer(dependency *common.Dependency) {
	initializeRoutes(dependency)
	srv := &http.Server{
		Addr:    ":3131",
		Handler: dependency.Router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
