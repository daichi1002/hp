package usecase

import (
	"backend/domain/repository"
	"net/http"

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

// 記事一覧取得
func (u *ArticleUsecase) ListArticles(c *gin.Context) {
	articles, err := u.articleRepository.ListArticles()

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}

// 記事詳細取得
func (u *ArticleUsecase) GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	article, err := u.articleRepository.GetArticle(id)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}
