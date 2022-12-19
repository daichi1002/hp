package server

import (
	"backend/domain/repository"
	"backend/usecase"
	"context"

	"github.com/gin-gonic/gin"
)

func NewApiServer(ctx context.Context, router *gin.Engine, repos repository.Repositories) {
	gitUsecase := usecase.NewGitUsecase(repos.GitRepository)

	router.GET("/languages", func(c *gin.Context) { gitUsecase.GetLanguages(c, ctx) })
}
