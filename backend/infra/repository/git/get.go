package git

import (
	"context"
	"fmt"

	"github.com/google/go-github/v48/github"
)

func (r *gitRepository) GetMostUsedLanguages(ctx context.Context) ([]map[string]int, error) {
	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := r.Repositories.List(ctx, "daichi1002", opt)

	if _, ok := err.(*github.RateLimitError); ok {
		return nil, fmt.Errorf("Githubとの連携に失敗しました。%v", err)
	}

	// Githubから各リポジトリの使用言語をAPIで取得
	var languages []map[string]int
	for _, repo := range repos {
		name := repo.GetName()
		res, _, err := r.Repositories.ListLanguages(ctx, "daichi1002", name)

		if err != nil {
			return nil, fmt.Errorf("Githubとの連携に失敗しました。%v", err)
		}

		languages = append(languages, res)
	}
	return languages, nil
}
