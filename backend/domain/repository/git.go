package repository

import (
	"context"

	"github.com/google/go-github/v48/github"
)

type GitRepository interface {
	GetMostUsedLanguages(ctx context.Context) ([]map[string]int, error)
	GetEvent(ctx context.Context) ([]*github.Event, error)
}
