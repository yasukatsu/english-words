package postgres

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"

	c "y_ara/english-words/server/internal/pkg/config"
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
		"dbname=%v host=%v port=%v user=%v password=%v sslmode=disable",
		c.Conf.Postgres.Db, c.Conf.Postgres.Host, c.Conf.Postgres.Port, c.Conf.Postgres.User, c.Conf.Postgres.Pass,
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
		log.Printf("Find err is %v", err)
		return nil, err
	}
	return words, nil
}
