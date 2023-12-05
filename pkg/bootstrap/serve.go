package bootstrap

import (
	"ginblog/pkg/config"
	"ginblog/pkg/html"
	"ginblog/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
