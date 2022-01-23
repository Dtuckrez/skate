package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

func open() *sql.DB {
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	db, err := sql.Open("postgres", "postgres://postgres:password@skate-postgres:5432/skate?sslmode=disable")
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println("All Okay did open")
	}

	return db
}

func GetBoards() {

	var db = open()
	err := db.Ping()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Ping okay")
	}

	rows, err := db.Query(`SELECT * FROM "boards"`)
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var manufacture string
		var uuid string
		var date string
		var mod sql.NullString
		err = rows.Scan(&uuid, &manufacture, &date, &mod)
		CheckError(err)

		fmt.Println(manufacture)
	}

	CheckError(err)

	defer db.Close()
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
