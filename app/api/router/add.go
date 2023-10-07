package router

import (
	"mpp/api/command"
	"mpp/api/omdb"
	"mpp/api/types"
	"mpp/error_util"
	"net/http"
)

func AddMovie(writer http.ResponseWriter, request *http.Request) {
	var movie types.Movie
	err := omdb.ReadJSONRequest(request.Body, &movie)
	error_util.CheckError(err)

	err = command.AddMovie(&movie)
	error_util.CheckError(err)

	err = writeMovieDetailsResponse(*movie.IMDbId, writer)
	error_util.CheckError(err)
}
