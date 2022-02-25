package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"skate/internal/api"
	"skate/internal/datastore/postgres"
	"skate/internal/queue"
)

func Start() {
	log.Println("API Running")

	var (
		port     = os.Getenv("PORT")
		dbURL    = os.Getenv("DB_URL")
		queueURL = os.Getenv("QUEUE_URL")
	)

	var (
		db = postgres.New(dbURL)
		q  = queue.NewAMQPQueue(queueURL, "skate")
	)

	if err := q.Open(func(err error) {
		log.Fatalln(err)
	}); err != nil {
		log.Fatalln(err)
	}
	mux := http.NewServeMux()

	mux.Handle("/", api.NewAPI(db, q))
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
