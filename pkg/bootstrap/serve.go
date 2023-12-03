package bootstrap

import (
	"ginblog/pkg/config"
	"ginblog/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
