package router

import (
	"mpp/api/command"
	"mpp/api/json_util"
	"mpp/error_util"
	"net/http"
)

func ListMovies(writer http.ResponseWriter, request *http.Request) {
	movies, err := command.GetMovieList()
	error_util.CheckError(err)
	err = json_util.WriteJSONResponse(writer, movies)
	error_util.CheckError(err)
}
