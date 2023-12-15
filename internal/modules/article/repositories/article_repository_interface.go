package repositories

import ArticleModel "ginblog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
}
