package controller

import (
	"url-shortener/internal/types"

	"github.com/gin-gonic/gin"

	"net/http"

	"url-shortener/internal/service"
)

type Controller struct {
	service *service.Service
}

// LongToShort receives a url and return its short representation.
// It expect a valid json request to deserialize to types.Url struct.
// It returns http status 200 and types.Url when everything went okay.
// It returns http status 422 and types.Message in deserialization errors.
func (controller Controller) LongToShort(c *gin.Context) {
	var longUrl types.Url

	if err := c.BindJSON(&longUrl); err != nil {
		c.JSON(http.StatusBadRequest, types.Message{Msg: "Deserialization Error"})
		return
	}

	var shortUrl string = controller.service.LongToShort(longUrl.Name)
	var response types.Url = types.Url{Name: shortUrl}
	c.JSON(http.StatusOK, response)
}

// ShortToLong receives a short url as uri parameter and return w/ a redirect to the long url previously registered.
// It returns an http response with status 307 (redirect to the long url)
func (controller Controller) ShortToLong(c *gin.Context) {
	var shortUrl string = c.Param("url")

	var longUrl string = controller.service.ShortToLong(shortUrl)

	c.Redirect(http.StatusTemporaryRedirect, longUrl)
}

// HealthCheck can be used to know if the server is running, like a health check endpoint.
// It returns http status 200 and a dummy types.Message if the server is up.
func (controller Controller) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, types.Message{Msg: "UP"})
}

func CreateController(service *service.Service) *Controller {
	return &Controller{service: service}
}
