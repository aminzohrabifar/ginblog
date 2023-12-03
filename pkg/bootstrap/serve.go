package bootstrap

import (
	"ginblog/pkg/config"
	"ginblog/pkg/routing"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Serve() {
	config.Set()
	routing.Init()
	router := routing.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app_name": viper.Get("App.Name"),
		})
	})
	routing.Serve()
}
