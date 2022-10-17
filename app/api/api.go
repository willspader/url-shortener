package api

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"url-shortener/app/types"
)

// LongUrlToShortUrl take a url and return its short representation.
// It expect a valid json request to deserialize to types.Url struct.
// It returns http status 200 and types.Url when everything went okay.
// It returns http status 422 and types.Message in deserialization errors.
func LongUrlToShortUrl(c *gin.Context) {
	var longUrl types.Url

	if err := c.BindJSON(&longUrl); err != nil {
		c.JSON(http.StatusBadRequest, types.Message{Msg: "Deserialization Error"})
		return
	}

	c.JSON(http.StatusOK, longUrl)
}

// CheckUpServer can be used to know if the server is running, like a health check endpoint.
// It returns http status 200 and a dummy types.Message if the server is up.
func CheckUpServer(c *gin.Context) {
	c.JSON(http.StatusOK, types.Message{Msg: "UP"})
}
