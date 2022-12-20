package server

import (
	"backend/domain/repository"
	"backend/usecase"
	"context"

	"github.com/gin-gonic/gin"
)

func NewApiServer(ctx context.Context, router *gin.Engine, repos repository.Repositories) {
	gitUsecase := usecase.NewGitUsecase(repos.GitRepository)

	router.GET("/commit_data", func(c *gin.Context) { gitUsecase.GetCommitData(c, ctx) })
	router.GET("/languages", func(c *gin.Context) { gitUsecase.GetLanguages(c) })
}
