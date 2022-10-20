package handlers

import (
	"url-shortener/internal/types"

	"github.com/gin-gonic/gin"

	"net/http"
)

// HealthCheck can be used to know if the server is running, like a health check endpoint.
// It returns http status 200 and a dummy types.Message if the server is up.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, types.Message{Msg: "UP"})
}
