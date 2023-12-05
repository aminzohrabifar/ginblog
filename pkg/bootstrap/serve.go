package bootstrap

import (
	"ginblog/pkg/config"
	"ginblog/pkg/html"
	"ginblog/pkg/routing"
	"ginblog/pkg/static"
)

func Serve() {
	config.Set()
	routing.Init()
	static.LoadStatic(routing.GetRouter())
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
