package repository

type GitRepository interface {
	GetMostUsedLanguages() (string, error)
}
