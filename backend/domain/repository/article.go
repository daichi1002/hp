package repository

import "backend/domain/model"

type ArticleRepository interface {
	ListArticles() ([]*model.Article, error)
	GetArticle(id string) (*model.Article, error)
}
