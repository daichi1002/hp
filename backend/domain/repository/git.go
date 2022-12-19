package repository

import "context"

type GitRepository interface {
	GetMostUsedLanguages(ctx context.Context) ([]map[string]int, error)
}
