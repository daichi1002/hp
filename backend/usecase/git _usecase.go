package usecase

import (
	"backend/domain/repository"
	"backend/util"
	"fmt"

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

func (u *GitUsecase) GetLanguages(c *gin.Context) {
	languages, err := u.gitRepository.GetMostUsedLanguages()

	fmt.Println(languages)
	fmt.Println(err)

}
