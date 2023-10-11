package router

import (
	"mpp/api/command"
	"mpp/api/json_util"
	"mpp/error_util"
	"net/http"
	"strconv"
)

const DEFAULT_PAGE = 0
const DEFAULT_PAGE_LIMIT = 50

func ListMovies(writer http.ResponseWriter, request *http.Request) {
	page, limit := getPaginationParameters(request)

	movies, err := command.GetMovieList(page, limit)
	error_util.CheckError(err)

	err = json_util.WriteJSONResponse(writer, movies)
	error_util.CheckError(err)
}

func getPaginationParameters(request *http.Request) (int, int) {
	pageString := request.URL.Query().Get("page")
	limitString := request.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = DEFAULT_PAGE
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		limit = DEFAULT_PAGE_LIMIT
	}

	return page, limit
}
