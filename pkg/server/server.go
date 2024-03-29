package server

import (
	"github.com/gin-gonic/gin"

	"url-shortener/internal/controller"
	"url-shortener/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartHttpServer(service *service.Service) {
	var controller *controller.Controller = controller.New(service)

	var router *gin.Engine = gin.Default()
	makeRoutes(router, controller)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func makeRoutes(router *gin.Engine, controller *controller.Controller) {
	router.GET("/", controller.HealthCheck)

	router.GET("/url/:url", controller.ShortToLong)

	router.POST("/url", controller.LongToShort)

	// API Documentation - Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
