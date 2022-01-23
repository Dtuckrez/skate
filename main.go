package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"v1/postgres"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Successfully connected!")

	log.Println("Running")

	handleRequests()
}

// which methods need to be called based on url
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getboards", getBoards)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

// servers index.html
func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page")
	log.Println(r.URL.Path)
	http.ServeFile(w, r, r.URL.Path[1:])
}

// return the boards that are currently stored in postgres
func getBoards(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		postgres.GetBoards()
	default:
		fmt.Fprintf(w, "Sorry, only GET")
	}

}
