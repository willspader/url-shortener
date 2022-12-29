package controller

import (
	"url-shortener/internal/types"

	"github.com/gin-gonic/gin"

	"net/http"

	"url-shortener/internal/service"

	_ "url-shortener/docs"
)

type Controller struct {
	service *service.Service
}

// @Summary receives a url and return its short representation.
// @ID long-to-short
// @Produce json
// @Success 200 {object} types.Url
// @Router /url [post]
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

// @Summary receives a short url as uri parameter and return w/ a redirect to the long url previously registered.
// @ID short-to-long
// @Param url path string true "url"
// @Success 307 {string} longUrl
// @Router /url/{url} [get]
func (controller Controller) ShortToLong(c *gin.Context) {
	var shortUrl string = c.Param("url")

	var longUrl string = controller.service.ShortToLong(shortUrl)

	c.Redirect(http.StatusTemporaryRedirect, longUrl)
}

// @Summary Can be used to know if the server is running.
// @ID health-check
// @Success 200 {object} types.Message
// @Router / [get]
func (controller Controller) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, types.Message{Msg: "UP"})
}

func New(service *service.Service) *Controller {
	return &Controller{service: service}
}
