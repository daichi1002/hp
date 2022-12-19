package server

import (
	"backend/domain/repository"
	"context"

	"github.com/gin-gonic/gin"
)

func NewApiServer(ctx context.Context, router *gin.Engine, repos repository.Repositories) {}
