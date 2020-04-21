package postgres

import "y_ara/english-words/backend/internal/pkg/db/postgres"

// ListWords ...
func ListWords() ([]postgres.Words, error) {
	err := postgres.NewClient()
	if err != nil {
		return nil, err
	}

	words, err := postgres.Cli.GetWords()
	if err != nil {
		return nil, err
	}
	return words, nil
}
