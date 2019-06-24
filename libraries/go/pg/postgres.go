package pg

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/jace-ys/super-smash-heroes/libraries/go/config"
)

type Conn struct {
	*sqlx.DB
}

type Config struct {
	Host          string
	Port          int
	User          string
	Password      string
	DbName        string
	DisableSSL    bool
	RetryInterval int
}

type Error struct {
	Code string
	Name string
}

func DefaultConfig() *Config {
	return &Config{
		Host:          config.Get("db.postgres.host").String("localhost"),
		Port:          config.Get("db.postgres.port").Int(5432),
		User:          "postgres",
		Password:      "mysecretpassword",
		DbName:        "postgres",
		DisableSSL:    true,
		RetryInterval: config.Get("db.postgres.retry.interval").Int(10),
	}
}

func Connect(cfg ...*Config) (*Conn, error) {
	cfg = append(cfg, DefaultConfig())
	c := cfg[0]
	var sslmode string
	switch c.DisableSSL {
	case true:
		sslmode = "disable"
	case false:
		sslmode = "verify-full"
	}
	psqlSourceInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s dbname=%s", c.Host, c.Port, c.User, c.Password, sslmode, c.DbName)
	for {
		db, err := sqlx.Open("postgres", psqlSourceInfo)
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err == nil {
			return &Conn{db}, nil
		}
		log.Printf("Retrying connection to database at address: %s:%d\n", c.Host, c.Port)
		time.Sleep(time.Second * time.Duration(c.RetryInterval))
	}
}

func (c *Conn) Close() error {
	return c.DB.Close()
}

func (c *Conn) Transact(fn func(*sqlx.Tx) error) error {
	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}
	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func GetError(err error) *Error {
	if e, ok := err.(*pq.Error); ok {
		return &Error{Code: string(e.Code), Name: e.Code.Name()}
	}
	return &Error{}
}
