package usecase

import (
	"backend/domain/repository"

	"github.com/gin-gonic/gin"
)

type ArticleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{
		articleRepository: repo,
	}
}

func (u *ArticleUsecase) ListArticles(c *gin.Context) {}

func (u *ArticleUsecase) GetArticle(c *gin.Context) {}
