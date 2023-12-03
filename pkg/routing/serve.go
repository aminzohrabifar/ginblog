package routing

import (
	"fmt"
	"ginblog/pkg/config"
	"log"
)

func Serve() {
	configs := config.Get()
	err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
