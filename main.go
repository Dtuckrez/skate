package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"v1/postgres"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "skate-postgres"
// 	port     = 8061
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "skate"
// )

func main() {

	fmt.Println("Successfully connected!")

	log.Println("Running")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/boards", func(w http.ResponseWriter, r *http.Request) {
		// 	fmt.Fprintf(w, "get me boards like init")
		// 	uuid := uuid.New()

		// 	// insert a row
		// 	sqlStatement := `INSERT INTO boards (id, manufacture)
		// 	VALUES ($1, $2)`
		// 	_, err = db.Exec(sqlStatement, uuid, "Birdhouse")
		// 	if err != nil {
		// 		panic(err)
		// 	} else {
		// 		fmt.Println("\nRow inserted successfully!")
		// 	}

		postgres.GetBoards()

	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
