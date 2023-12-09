package repositories

import articleModels "ginblog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	list(limit int) []articleModels.Article
}
