package psql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jace-ys/super-smash-heroes/libraries/go/config"
	"github.com/jace-ys/super-smash-heroes/libraries/go/errors"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

var (
	driver = "postgres"
	table  = config.Get("db.superhero.table").String("superheroes")
)

type Client struct {
	db *sql.DB
}

func Open(dataSource string) (*Client, error) {
	for {
		db, err := sql.Open(driver, dataSource)
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err == nil {
			return &Client{db}, nil
		}
		log.Println("Retrying connection to database")
		retryInterval := config.Get("db.superhero.retry.interval").Int(5)
		time.Sleep(time.Second * time.Duration(retryInterval))
	}
}

func (psql *Client) Close() error {
	return psql.db.Close()
}

func (c *Client) GetAll() ([]*pb.SuperheroResponse, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := c.db.Query(query)
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

func (c *Client) FindById(id int32) (*pb.SuperheroResponse, error) {
	var row pb.SuperheroResponse
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", table)
	rows := c.db.QueryRow(query, id)
	err := rows.Scan(&row.Id, &row.FullName, &row.AlterEgo, &row.ImageUrl)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.SuperheroNotFound
	case err != nil:
		return nil, err
	}
	return &row, nil
}

func (c *Client) Find(key, val string) (*pb.SuperheroResponse, error) {
	var row pb.SuperheroResponse
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", table, key)
	rows := c.db.QueryRow(query, val)
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
	found, _ := c.Find("full_name", s.FullName)
	if found != nil {
		return -1, errors.SuperheroExists
	}
	var id int32
	query := fmt.Sprintf("INSERT INTO %s (full_name, alter_ego, image_url) VALUES ($1, $2, $3) RETURNING id", table)
	err := c.db.QueryRow(query, s.FullName, s.AlterEgo, s.ImageUrl).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (psql *Client) Delete(id int32) error {
	statement := fmt.Sprintf("DELETE FROM %s WHERE id=$1", table)
	_, err := psql.db.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}
