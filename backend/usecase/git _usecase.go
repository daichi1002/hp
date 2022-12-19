package usecase

import (
	"backend/domain/model"
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
	usedlanguages, err := u.gitRepository.GetMostUsedLanguages(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	response := AddLanguages(usedlanguages)

	c.JSON(http.StatusOK, response)
}

func AddLanguages(usedlanguages []map[string]int) []*model.Language {
	totalLanguage := make(map[string]int)
	for _, language := range usedlanguages {
		if language == nil {
			continue
		}
		for key, value := range language {
			if _, ok := totalLanguage[key]; ok {
				totalLanguage[key] = totalLanguage[key] + value
			} else {
				totalLanguage[key] = value
			}
		}
	}

	response := []*model.Language{}
	for key, value := range totalLanguage {
		element := &model.Language{
			Name:  key,
			Ratio: value,
		}
		response = append(response, element)
	}

	return response
}
