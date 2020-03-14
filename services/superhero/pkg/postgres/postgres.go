package postgres

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Client interface {
	Transact(ctx context.Context, fn func(*sqlx.Tx) error) error
	Close() error
}

type Config struct {
	Host     string
	User     string
	Password string
	Database string
}

type postgresClient struct {
	config *Config
	*sqlx.DB
}

func NewPostgresClient(host, user, password, database string) (*postgresClient, error) {
	client := postgresClient{
		config: &Config{
			Host:     host,
			User:     user,
			Password: password,
			Database: database,
		},
	}

	if err := client.init(); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *postgresClient) init() error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.config.User,
		c.config.Password,
		c.config.Host,
		c.config.Database,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("postgres connection failed: %w", err)
	}
	db.MapperFunc(toLowerSnakeCase)

	c.DB = db
	return nil
}

func toLowerSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (c *postgresClient) Transact(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := c.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("postgres transaction failed: %w", err)
	}

	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("postgres transaction failed: %w", err)
	}

	return tx.Commit()
}
