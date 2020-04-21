package postgres

import (
	"fmt"

	c "y_ara/english-words/backend/internal/pkg/config"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// Cli is ...
var Cli *PostgresClient

// PostgresClient ...
type PostgresClient struct {
	db *xorm.Engine
}

// NewClient is ...
func NewClient() error {
	src := fmt.Sprintf(
		"dbname=%v port=%v user=%v password=%v sslmode=disable",
		c.Conf.Postgres.Db, c.Conf.Postgres.Port, c.Conf.Postgres.User, c.Conf.Postgres.Pass,
	)
	engine, err := xorm.NewEngine("postgres", src)
	if err != nil {
		return err
	}

	Cli = &PostgresClient{db: engine}
	return nil
}

// GetWords ...
func (c *PostgresClient) GetWords() ([]Words, error) {
	defer c.db.Close()
	var words []Words
	if err := c.db.Find(&words); err != nil {
		return nil, err
	}
	return words, nil
}
