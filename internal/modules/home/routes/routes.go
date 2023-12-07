package routes

import (
	"ginblog/internal/modules/home/controllers"
	"ginblog/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	homeController := controllers.New()
	router.GET("/", homeController.Index)

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page",
		})
	})
}
