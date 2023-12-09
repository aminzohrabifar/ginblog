package repositories

import (
	articleModels "ginblog/internal/modules/article/models"
	"ginblog/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{DB: database.Connection()}
}
func (ArticleRepository *ArticleRepository) list(limit int) []articleModels.Article {
	var articles []articleModels.Article
	ArticleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}
