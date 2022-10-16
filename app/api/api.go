package api

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"url-shortener/app/types"
)

func LongUrlToShortUrl(c *gin.Context) {
	var longUrl types.Url

	if err := c.BindJSON(&longUrl); err != nil {
		return // TODO
	}

	c.JSON(http.StatusOK, longUrl)
}

func CheckUpServer(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})

}
