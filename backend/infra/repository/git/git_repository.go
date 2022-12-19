package git

import (
	"backend/domain/repository"

	"github.com/google/go-github/v48/github"
)

type gitRepository struct {
	*github.Client
}

func NewGitRepository(client *github.Client) repository.GitRepository {
	return &gitRepository{
		client,
	}
}
