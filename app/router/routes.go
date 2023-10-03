package router

import (
	"fmt"
	"net/http"
	"os"
)

const MOVIES_ENDPOINT string = "/movies"
const MOVIE_ENDPOINT string = "/movies/:id"

func StartAPI() {
	address := getAddress()

	registerRoutes()

	http.ListenAndServe(address, nil)
}

func registerRoutes() {
	addEndpoint(MOVIES_ENDPOINT, moviesHandler)
	addEndpoint(MOVIE_ENDPOINT, movieHandler)
}

func moviesHandler(w http.ResponseWriter, request *http.Request) {
	writer := &w
	switch request.Method {
	case http.MethodGet:
		OnMovieListRequest(writer, request)
	case http.MethodPost:
		OnMovieAddRequest(writer, request)
	default:
		methodNotAllowed(writer)
	}
}

func movieHandler(w http.ResponseWriter, request *http.Request) {
	writer := &w
	switch request.Method {
	case http.MethodGet:
		OnMovieDetailsRequest(writer, request)
	case http.MethodDelete:
		OnMovieDeleteRequest(writer, request)
	default:
		methodNotAllowed(writer)
	}
}

func addEndpoint(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(endpoint, handler)
}

func methodNotAllowed(writer *http.ResponseWriter) {
	(*writer).WriteHeader(http.StatusMethodNotAllowed)
}

func getAddress() string {
	port := 8090
	hostname := os.Getenv("API_HOST")
	if hostname == "" {
		hostname = "localhost"
	}

	return fmt.Sprintf("%s:%d", hostname, port)
}
