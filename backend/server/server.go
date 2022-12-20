package server

import (
	"backend/domain/repository"
	"backend/usecase"
	"context"

	"github.com/gin-gonic/gin"
)

func NewApiServer(ctx context.Context, router *gin.Engine, repos repository.Repositories) {
	gitUsecase := usecase.NewGitUsecase(repos.GitRepository)

	router.GET("/language_data", func(c *gin.Context) { gitUsecase.FetchLanguageData(c, ctx) })
	router.GET("/languages", func(c *gin.Context) { gitUsecase.GetLanguages(c) })
	router.GET("/commit_data", func(c *gin.Context) { gitUsecase.FetchCommitData(c, ctx) })
	router.GET("/contributions", func(c *gin.Context) { gitUsecase.GetContributions(c) })
}
