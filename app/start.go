package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"net/http"
)

type url struct {
	Url string `json:"url"`
}

func main() {
	router := gin.Default()

	router.GET("/", checkUpServer)

	router.POST("/url", longUrlToShortUrl)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func longUrlToShortUrl(c *gin.Context) {
	var longUrl url

	if err := c.BindJSON(&longUrl); err != nil {
		return // TODO
	}

	fmt.Println(longUrl)

	c.JSON(http.StatusOK, longUrl)
}

func checkUpServer(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})

}
