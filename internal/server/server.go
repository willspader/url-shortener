package server

import (
	"github.com/gin-gonic/gin"

	"url-shortener/internal/handlers"
)

func StartHttpServer() {
	var router *gin.Engine = gin.Default()

	makeRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func makeRoutes(router *gin.Engine) {
	router.GET("/", handlers.HealthCheck)

	router.POST("/url", handlers.LongToShort)
}
