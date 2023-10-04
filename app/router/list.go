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

	// get rid of the plot field
	for _, movie := range movies {
		movie.Plot_summary = nil
	}

	err = omdb.WriteJSONResponse(writer, movies)
	error_util.CheckError(err)
}
