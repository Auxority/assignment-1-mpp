package router

import (
	"mpp/command"
	"mpp/error_util"
	"mpp/omdb"
	"mpp/types"
	"net/http"
)

func AddMovie(writer http.ResponseWriter, request *http.Request) {
	var movie types.Movie
	err := omdb.ReadJSONRequest(request.Body, &movie)
	error_util.CheckError(err)

	command.AddMovie(&movie)
	writeMovieDetailsResponse(*movie.IMDbId, writer)
}
