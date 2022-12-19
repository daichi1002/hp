package usecase

import (
	"backend/domain/repository"
	"backend/util"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logger = util.NewLogger()

type GitUsecase struct {
	gitRepository repository.GitRepository
}

func NewGitUsecase(repo repository.GitRepository) *GitUsecase {
	return &GitUsecase{
		gitRepository: repo,
	}
}

func (u *GitUsecase) GetLanguages(c *gin.Context, ctx context.Context) {
	languages, err := u.gitRepository.GetMostUsedLanguages(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, languages)
}
