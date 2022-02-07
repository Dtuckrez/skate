package postgres

import (
	"database/sql"
	"skate/internal/datastore"
)

type PSQL struct {
	conn string
}

func New(conn string) PSQL {
	return PSQL{conn}
}

func (i PSQL) List() ([]datastore.Board, error) {
	db, err := sql.Open("postgres", i.conn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT manufacture FROM "boards"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boards []datastore.Board

	// checks each row of the sql query and creates a board struct and adds to an array
	for rows.Next() {
		var b datastore.Board
		if err := rows.Scan(&b.Manufacture); err != nil {
			return nil, err
		}
		boards = append(boards, b)
	}
	return boards, nil
	// return []datastore.Board{}, nil
}
