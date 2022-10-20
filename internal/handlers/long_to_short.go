package handlers

import (
	"url-shortener/internal/types"

	"github.com/gin-gonic/gin"

	"net/http"
)

// LongToShort take a url and return its short representation.
// It expect a valid json request to deserialize to types.Url struct.
// It returns http status 200 and types.Url when everything went okay.
// It returns http status 422 and types.Message in deserialization errors.
func LongToShort(c *gin.Context) {
	var longUrl types.Url

	if err := c.BindJSON(&longUrl); err != nil {
		c.JSON(http.StatusBadRequest, types.Message{Msg: "Deserialization Error"})
		return
	}

	c.JSON(http.StatusOK, longUrl)
}
