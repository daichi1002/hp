package main

import (
	"backend/infra"
	"backend/util"
	"context"
	"time"

	domainrepo "backend/domain/repository"
	"backend/infra/repository/article"
	"backend/infra/repository/git"
	"backend/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var logger = util.NewLogger()

func main() {
	// 環境変数読み込み
	err := godotenv.Load(".env")

	if err != nil {
		logger.Fatalf("Failed to connect env: %v", err)
	}

	ctx := context.Background()
	// DB初期化処理
	gormHandler := infra.NewGormHandler()
	githubClient := infra.NewGithubApiClient(ctx)

	repos := domainrepo.Repositories{
		ArticleRepository: article.NewArticleRepository(gormHandler),
		GitRepository:     git.NewGitRepository(githubClient),
	}

	router := gin.Default()

	setCors(router)

	server.NewApiServer(ctx, router, repos)

	router.Run()
	logger.Info("server start!")
}

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
