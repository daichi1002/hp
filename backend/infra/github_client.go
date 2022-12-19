package infra

import (
	"context"
	"os"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func NewGithubApiClient(ctx context.Context) *github.Client {
	token := os.Getenv("ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}
