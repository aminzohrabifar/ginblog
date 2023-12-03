package cmd

import (
	"fmt"
	"ginblog/pkg/config"
	"github.com/gin-gonic/gin"
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
	configs := config.Get()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app_name": configs.App.Name,
		})
	})
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println(err)
	}
}
