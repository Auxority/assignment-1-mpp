package router

import "net/http"

func OnMovieListRequest(writer *http.ResponseWriter, request *http.Request) {
	(*writer).Write([]byte("LIST DA MOVIESSS"))
	// var movies []*types.Movie
	// queryMovies(&movies)

	// context.IndentedJSON(http.StatusOK, &movies)
}
