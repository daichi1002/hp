package server

import (
	"backend/domain/repository"
	"backend/usecase"
	"context"

	"github.com/gin-gonic/gin"
)

func NewApiServer(ctx context.Context, router *gin.Engine, repos repository.Repositories) {
	gitUsecase := usecase.NewGitUsecase(repos.GitRepository)
	articleUsecase := usecase.NewArticleUsecase(repos.ArticleRepository)

	// GithubAPIによるデータ取得
	router.GET("/language_data", func(c *gin.Context) { gitUsecase.FetchLanguageData(c, ctx) })
	router.GET("/commit_data", func(c *gin.Context) { gitUsecase.FetchCommitData(c, ctx) })
	// Githubに関するデータのレスポンス
	router.GET("/languages", func(c *gin.Context) { gitUsecase.GetLanguages(c) })
	router.GET("/contributions", func(c *gin.Context) { gitUsecase.GetContributions(c) })
	// 記事に関するデータのレスポンス
	router.GET("/articles", func(c *gin.Context) { articleUsecase.ListArticles(c) })
	router.GET("/article/:id", func(c *gin.Context) { articleUsecase.GetArticle(c) })
}
