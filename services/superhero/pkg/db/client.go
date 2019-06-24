package db

import (
	"database/sql"

	"github.com/jace-ys/super-smash-heroes/libraries/go/config"
	"github.com/jace-ys/super-smash-heroes/libraries/go/errors"
	"github.com/jace-ys/super-smash-heroes/libraries/go/pg"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

type Client struct {
	db *pg.Conn
}

func NewClient() (*Client, error) {
	pgConfig := &pg.Config{
		Host:          config.Get("db.superhero.host").String("localhost"),
		Port:          config.Get("db.superhero.port").Int(5432),
		User:          "postgres",
		Password:      "mysecretpassword",
		DbName:        "postgres",
		DisableSSL:    true,
		RetryInterval: config.Get("db.superhero.retry.interval").Int(10),
	}
	conn, err := pg.Connect(pgConfig)
	if err != nil {
		return nil, err
	}
	return &Client{conn}, nil
}

func (c *Client) Teardown() {
	c.db.Close()
}

func (c *Client) GetAll() ([]*pb.SuperheroResponse, error) {
	rows, err := c.db.Query("SELECT * FROM superheroes")
	if err != nil {
		return nil, err
	}
	var data []*pb.SuperheroResponse
	for rows.Next() {
		var row pb.SuperheroResponse
		if err = rows.Scan(&row.Id, &row.FullName, &row.AlterEgo, &row.ImageUrl); err != nil {
			return nil, err
		}
		data = append(data, &row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) FindByID(id int32) (*pb.SuperheroResponse, error) {
	var row pb.SuperheroResponse
	rows := c.db.QueryRow("SELECT * FROM superheroes WHERE id=$1", id)
	err := rows.Scan(&row.Id, &row.FullName, &row.AlterEgo, &row.ImageUrl)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.SuperheroNotFound
	case err != nil:
		return nil, err
	}
	return &row, nil
}

func (c *Client) Insert(s *pb.SuperheroResponse) (int32, error) {
	var id int32
	err := c.db.QueryRow("INSERT INTO superheroes (full_name, alter_ego, image_url) VALUES ($1, $2, $3) RETURNING id", s.FullName, s.AlterEgo, s.ImageUrl).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (c *Client) DeleteByID(id int32) error {
	_, err := c.db.Exec("DELETE FROM superheroes WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
