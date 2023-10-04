package omdb

import (
	"fmt"
	"mpp/types"
	"net/http"
)

const API_KEY = "899f543d"

func buildUrl(id *string) string {
	return fmt.Sprintf(
		"http://www.omdbapi.com/?apikey=%s&i=%s",
		API_KEY,
		*id,
	)
}

func getRequest(url *string, data any) error {
	response, err := http.Get(*url)
	if err != nil {
		return fmt.Errorf("getRequest: failed to execute get request: %w", err)
	}

	err = ReadJSONRequest(response.Body, data)
	if err != nil {
		return fmt.Errorf("getRequest: %w", err)
	}

	return nil
}

func GetMovieSummary(id *string, summary *types.MovieSummary) error {
	requestUrl := buildUrl(id)
	err := getRequest(&requestUrl, &summary)
	if err != nil {
		return fmt.Errorf("GetMovieSummary: %w", err)
	}

	if summary.OK == "False" {
		return fmt.Errorf("GetMovieSummary: the API returned an error: %s", summary.Error)
	}

	return nil
}
