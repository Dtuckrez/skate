package api

// type MyInterface interface {
// 	myFunction()
// }

// type MyStruct struct {

// }
// func(i MyStruct) myFunction(){

// }

import (
	"encoding/json"
	"net/http"
	"skate/internal/datastore"
	"skate/internal/queue"
	"strings"
)

type API struct {
	routes    []Route
	datastore datastore.BoardReader
	queue     queue.Queue
}

// NewAPI represents the function to create a new API object for the supplied input.
func NewAPI(datastore datastore.BoardReader, queue queue.Queue) http.Handler {
	out := API{
		datastore: datastore,
		queue:     queue,
	}

	out.routes = []Route{
		NewRoute(http.MethodGet, "/getBoards", out.getBoards()),
		NewRoute(http.MethodPost, "/setBoard", out.setBoard()),
	}

	return out
}

// ServeHTTP represents the function to serve HTTP requests. This function is the implemtation for the http.Handler interface.
func (i API) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	handler := http.NotFoundHandler()

	for _, route := range i.routes {
		if r.Method == route.Method && strings.EqualFold(r.URL.Path, route.Path) {
			handler = route.Handler
		}
	}
	handler.ServeHTTP(w, r)
}

func (i API) getBoards() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boards, err := i.datastore.List()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		buff, err := json.Marshal(boards)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Write(buff)
		// w.Write([]byte("Hello World"))
	}
}

func (i API) setBoard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := i.queue.Enqueue([]byte("QUEUE")); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	}
}

// Route represents the type for an API route.
type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

// NewRoute represents the function to create a new Route object for the supplied input.
func NewRoute(method, path string, handler http.Handler) Route {
	return Route{method, path, handler}
}
