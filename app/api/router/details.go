package router

import (
	"fmt"
	"mpp/api/command"
	"mpp/api/json_util"
	"mpp/error_util"
	"net/http"
)

func MovieDetails(writer http.ResponseWriter, request *http.Request) {
	id := GetUrlId(request)
	err := writeMovieDetailsResponse(id, writer)
	if err != nil {
		error_util.CheckError(err)
		writer.WriteHeader(http.StatusNotFound)
	}
}

func writeMovieDetailsResponse(id string, writer http.ResponseWriter) error {
	movie, err := command.GetMovieDetails(&id)
	if err != nil {
		return fmt.Errorf("writeMovieDetailsResponse: %w", err)
	}

	err = json_util.WriteJSONResponse(writer, movie)
	if err != nil {
		return fmt.Errorf("writeMovieDetailsResponse: %w", err)
	}

	return nil
}
