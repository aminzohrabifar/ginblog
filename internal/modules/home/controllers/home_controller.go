package controllers

import (
	ArticleRepository "ginblog/internal/modules/article/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{articleRepository: ArticleRepository.New()}
}
func (controller *Controller) Index(c *gin.Context) {
	//html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home Page",
	//})
	c.JSON(http.StatusOK, gin.H{"articles": controller.articleRepository.List(2)})
}
