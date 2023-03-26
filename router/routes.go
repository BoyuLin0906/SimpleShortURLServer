package router

import (
	"net/http"
	"simple_short_url_server/controllers"

	"github.com/gin-gonic/gin"
)

type ShortenURL struct {
	FullUrl string `json:"url"`
}

func shorten_url(c *gin.Context) {
	var short_url_data ShortenURL
	err := c.Bind(&short_url_data)
	if err != nil {
		return
	}

	short_url := controllers.SetShortURL(short_url_data.FullUrl)
	c.JSON(http.StatusOK, gin.H{
		"short_url": short_url,
	})
}

func redirect_to_url(c *gin.Context) {
	short_url := c.Param("short_url")
	original_url := controllers.GetOriginalURL(short_url)
	if original_url == "" {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
	c.Redirect(http.StatusFound, original_url)
}

func InitServer() *gin.Engine {
	router := gin.Default()
	URLGroup := router.Group("/url")
	{
		URLGroup.GET("/:short_url", redirect_to_url)
		URLGroup.POST("/shorten", shorten_url)
	}

	return router
}
