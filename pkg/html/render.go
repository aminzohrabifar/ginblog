package html

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Render(c *gin.Context, status int, name string, data gin.H) {
	data["APP_NAME"] = viper.Get("App.Name")
	c.HTML(status, name, data)
}
