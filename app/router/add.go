package router

import "net/http"

func OnMovieAddRequest(writer *http.ResponseWriter, request *http.Request) {
	(*writer).Write([]byte("ADD A MOVIE"))
	// var movie types.Movie

	// err := context.BindJSON(&movie)
	// error_util.CheckError(err)

	// command.AddMovie(&movie)

	// context.IndentedJSON(http.StatusOK, &movie)
}
