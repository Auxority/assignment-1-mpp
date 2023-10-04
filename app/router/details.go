package router

import (
	"fmt"
	"mpp/command"
	"mpp/omdb"
	"net/http"
)

func MovieDetails(writer http.ResponseWriter, request *http.Request) {
	id := GetUrlId(request)
	err := writeMovieDetailsResponse(id, writer)
	if err != nil {
		fmt.Printf("Could not find movie! - %s\n", err.Error())
		writer.WriteHeader(http.StatusNotFound)
	}
}

func writeMovieDetailsResponse(id string, writer http.ResponseWriter) error {
	movie, err := command.GetMovieDetails(&id)
	if err != nil {
		return fmt.Errorf("writeMovieDetailsResponse: %w", err)
	}

	err = omdb.WriteJSONResponse(writer, movie)
	if err != nil {
		return fmt.Errorf("writeMovieDetailsResponse: %w", err)
	}

	return nil
}
