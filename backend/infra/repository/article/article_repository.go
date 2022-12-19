package article

import (
	"backend/domain/repository"
	"backend/infra"
)

type articleRepository struct {
	*infra.GormHandler
}

// NewArticleRepository NewArticleRepositoryを生成します。.
func NewArticleRepository(gormHandler *infra.GormHandler) repository.ArticleRepository {
	return &articleRepository{
		gormHandler,
	}
}
