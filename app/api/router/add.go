package router

import (
	"fmt"
	"mpp/api/command"
	"mpp/api/json_util"
	"mpp/api/omdb"
	"mpp/api/types"
	"mpp/error_util"
	"net/http"
)

const DEFAULT_RATING = 0.0

func AddMovie(writer http.ResponseWriter, request *http.Request) {
	var movie *types.Movie
	err := json_util.ReadJSONRequest(request.Body, &movie)
	if err != nil {
		movie, err = searchMovie(writer, request, movie.IMDbId)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	}

	err = command.AddMovie(movie)
	error_util.CheckError(err)

	err = writeMovieDetailsResponse(*movie.IMDbId, writer)
	error_util.CheckError(err)
}

func searchMovie(writer http.ResponseWriter, request *http.Request, id *string) (*types.Movie, error) {
	details, err := omdb.GetMovieDetails(id)
	if err != nil {
		return nil, fmt.Errorf("searchMovie: %w", err)
	}

	releaseYear, err := json_util.ParseInteger(&details.ReleaseYear)
	if err != nil {
		return nil, fmt.Errorf("searchMovie: %w", err)
	}

	rating, err := json_util.ParseFloat(&details.IMDbRating)
	if err != nil {
		testRating := DEFAULT_RATING
		rating = &testRating
	}

	movie := types.Movie{
		IMDbId:       &details.IMDbId,
		Title:        &details.Title,
		IMDbRating:   rating,
		ReleaseYear:  releaseYear,
		Plot_summary: &details.Plot_summary,
	}

	return &movie, nil
}
