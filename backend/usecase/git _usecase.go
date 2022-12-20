package usecase

import (
	"backend/constant"
	"backend/domain/model"
	"backend/domain/repository"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v48/github"

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

func (u *GitUsecase) FetchCommitData(c *gin.Context, ctx context.Context) {
	events, err := u.gitRepository.GetEvent(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	response := AddCommits(events)

	// JSONファイルに書き込み
	file, err := os.Create("./contents/events.json")
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

func AddCommits(events []*github.Event) []*model.Commit {
	commit := make(map[string]int)
	for _, event := range events {
		if event == nil {
			continue
		}
		formatDate := event.CreatedAt.Format(constant.YYYYMMDD)

		if _, ok := commit[formatDate]; ok {
			commit[formatDate]++
		} else {
			commit[formatDate] = 0
		}
	}

	response := []*model.Commit{}
	for key, value := range commit {
		element := &model.Commit{
			Date:  key,
			Count: value,
		}
		response = append(response, element)
	}

	return response
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

	var languages []*model.Language
	res, err := getJsonData(file, languages)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
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

	var commits []*model.Commit
	res, err := getJsonData(file, commits)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func getJsonData[T model.Git](file *os.File, domain T) (T, error) {
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&domain); err != nil {
		return nil, fmt.Errorf("JSONデコードに失敗しました。%v", err)
	}

	return domain, nil
}
