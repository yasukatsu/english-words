package postgres

import (
	"log"
	"y_ara/english-words/backend/internal/pkg/db/postgres"
)

// ListWords ...
func ListWords() ([]postgres.Words, error) {
	err := postgres.NewClient()
	if err != nil {
		log.Printf("NewClient err is %v", err)
		return nil, err
	}

	words, err := postgres.Cli.GetWords()
	if err != nil {
		log.Printf("GetWords err is %v", err)
		return nil, err
	}
	return words, nil
}
