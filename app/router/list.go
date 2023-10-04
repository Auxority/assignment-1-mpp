package router

import (
	"mpp/command"
	"mpp/error_util"
	"mpp/omdb"
	"net/http"
)

func ListMovies(writer http.ResponseWriter, request *http.Request) {
	movies, err := command.GetMovieList()
	error_util.CheckError(err)

	err = omdb.WriteJSONResponse(writer, movies)
	error_util.CheckError(err)
}
