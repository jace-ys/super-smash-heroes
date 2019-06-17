package psql

import (
	"database/sql"
	"fmt"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

const (
	driver = "postgres"
	table  = "superheroes"
)

type Client struct {
	db *sql.DB
}

func Open(dataSource string) (*Client, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Client{db}, nil
}

func (psql *Client) Close() error {
	return psql.db.Close()
}

// func (psql *psqlClient) Insert(table, number string) (int, error) {
// 	var id int
// 	query := fmt.Sprintf("INSERT INTO %s (num) VALUES ($1) RETURNING id", table)
// 	err := psql.db.QueryRow(query, number).Scan(&id)
// 	if err != nil {
// 		return -1, err
// 	}
// 	return id, nil
// }
//
// func (psql *psqlClient) Find(table string, number string) (*phone.Number, error) {
// 	var row phone.Number
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE num=$1", table)
// 	rows := psql.db.QueryRow(query, number)
// 	err := rows.Scan(&row.ID, &row.Number)
// 	if err != nil {
// 		switch {
// 		case err == sql.ErrNoRows:
// 			return nil, nil
// 		default:
// 			return nil, err
// 		}
// 	}
// 	return &row, nil
// }
//
// func (psql *psqlClient) Update(table string, p *phone.Number) error {
// 	statement := fmt.Sprintf("UPDATE %s SET num=$1 WHERE id=$2", table)
// 	_, err := psql.db.Exec(statement, p.Number, p.ID)
// 	return err
// }
//
// func (psql *psqlClient) Delete(table string, id int) error {
// 	statement := fmt.Sprintf("DELETE FROM %s WHERE id=$1", table)
// 	_, err := psql.db.Exec(statement, id)
// 	return err
// }

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
