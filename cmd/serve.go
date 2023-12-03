package cmd

import (
	"ginblog/pkg/config"
	"ginblog/pkg/routing"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve app on dev server",
	Long:  `application will be served on host and port defined in config.yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
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
