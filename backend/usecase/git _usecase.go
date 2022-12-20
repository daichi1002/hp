package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/util"
	"context"
	"encoding/json"
	"net/http"
	"os"

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

// GithubAPIを使用して、データ取得後JSONファイルに書き込み
func (u *GitUsecase) FetchLanguageData(c *gin.Context, ctx context.Context) {
	usedlanguages, err := u.gitRepository.GetMostUsedLanguages(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	response := AddLanguages(usedlanguages)

	// JSONファイルに書き込み
	file, err := os.Create("./contents/language.json")
	if err != nil {
		logger.Error(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(response); err != nil {
		logger.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (u *GitUsecase) FetchCommitData(c *gin.Context, ctx context.Context) {
	events, err := u.gitRepository.GetEvent(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	// JSONファイルに書き込み
	file, err := os.Create("./contents/events.json")
	if err != nil {
		logger.Error(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(events); err != nil {
		logger.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// JSONファイルから読み込み、レスポンスを返す
func (u *GitUsecase) GetLanguages(c *gin.Context) {
	file, err := os.Open("./contents/language.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	defer file.Close()

	// JSON ファイルの内容を Person 構造体データとして読み出す
	var languages []*model.Language
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&languages); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, languages)
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

// JSONファイルから読み込み、レスポンスを返す
func (u *GitUsecase) GetContributions(c *gin.Context) {
	file, err := os.Open("./contents/events.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	defer file.Close()

	// JSON ファイルの内容を Person 構造体データとして読み出す
	var languages []*model.Language
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&languages); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, languages)
}
