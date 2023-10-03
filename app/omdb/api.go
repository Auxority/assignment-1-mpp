package omdb

import (
	"encoding/json"
	"fmt"
	"mpp/error_util"
	"net/http"
)

type MovieSummary struct {
	Response string
	Error    string
}

const API_KEY = ""

func buildUrl(id *string) string {
	return fmt.Sprintf(
		"http://www.omdbapi.com/?apikey=%s&i=%s",
		API_KEY,
		*id,
	)
}

func parseJSON(res *http.Response, data any) {
	err := json.NewDecoder(res.Body).Decode(data)
	error_util.CheckError(err)
}

func getRequest(url *string, data any) {
	res, err := http.Get(*url)
	error_util.CheckError(err)
	parseJSON(res, data)
}

func GetMovieSummary(id *string, summary *MovieSummary) {
	requestUrl := buildUrl(id)
	go getRequest(&requestUrl, &summary)
}
