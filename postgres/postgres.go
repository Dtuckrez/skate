package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

type board struct {
	manufacture string
}

// Opens a connection to the postgres DB using some simple values
func openDBconnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:password@skate-postgres:5432/skate?sslmode=disable")
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println("All Okay did open")
	}

	return db
}

// calls a Select SQL query to fetch boards from Postgres
func GetBoards() {

	var db = openDBconnection()
	err := db.Ping()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Ping okay")
	}

	rows, err := db.Query(`SELECT manufacture FROM "boards"`)
	CheckError(err)

	var boards []board

	defer rows.Close()

	// checks each row of the sql query and creates a board struct and adds to an array
	for rows.Next() {
		var b board

		err = rows.Scan(&b.manufacture)
		boards = append(boards, b)
		CheckError(err)

		fmt.Println(boards)
	}

	CheckError(err)

	defer db.Close()
}

// Error handler to print out the error longs
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
