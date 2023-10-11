package omdb

import (
	"fmt"
	"mpp/api/json_util"
	"mpp/api/types"
	"net/http"
)

const API_KEY = "899f543d"

func buildUrl(id *string) string {
	return fmt.Sprintf(
		"https://www.omdbapi.com/?apikey=%s&i=%s",
		API_KEY,
		*id,
	)
}

func getRequest(url *string, data any) error {
	response, err := http.Get(*url)
	if err != nil {
		return fmt.Errorf("getRequest: failed to execute get request: %w", err)
	}

	err = json_util.ReadJSONRequest(response.Body, data)
	if err != nil {
		return fmt.Errorf("getRequest: %w", err)
	}

	return nil
}

func GetMovieDetails(id *string) (*types.MovieDetails, error) {
	var details types.MovieDetails
	requestUrl := buildUrl(id)
	err := getRequest(&requestUrl, &details)
	if err != nil {
		return nil, fmt.Errorf("GetMovieDetails: %w", err)
	}

	if details.OK == "False" || details.OK == "" {
		return nil, fmt.Errorf("GetMovieDetails: the API returned an error: %s", details.Error)
	}

	return &details, nil
}
