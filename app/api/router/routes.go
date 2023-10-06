package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartAPI() {
	address := getAddress()
	registerRoutes()
	http.ListenAndServe(address, nil)
}

func registerRoutes() {
	const MOVIES_API_ENDPOINT string = "/movies"
	const MOVIE_API_ENDPOINT string = "/movies/{id:tt[0-9]+}"

	router := mux.NewRouter()

	router.HandleFunc(MOVIES_API_ENDPOINT, ListMovies).Methods(http.MethodGet)
	router.HandleFunc(MOVIE_API_ENDPOINT, MovieDetails).Methods(http.MethodGet)
	router.HandleFunc(MOVIES_API_ENDPOINT, AddMovie).Methods(http.MethodPost)
	router.HandleFunc(MOVIE_API_ENDPOINT, DeleteMovie).Methods(http.MethodDelete)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	http.Handle("/", router)
}

func getAddress() string {
	port := 8090
	hostname := os.Getenv("API_HOST")
	if hostname == "" {
		hostname = "localhost"
	}

	return fmt.Sprintf("%s:%d", hostname, port)
}
