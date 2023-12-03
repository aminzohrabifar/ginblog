package routing

import (
	"ginblog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}
func GetRouter() *gin.Engine {
	return router
}
func RegisterRoutes() {
	routes.Routes(GetRouter())
}
