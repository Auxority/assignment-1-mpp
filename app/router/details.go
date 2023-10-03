package router

import "net/http"

func OnMovieDetailsRequest(writer *http.ResponseWriter, request *http.Request) {
	(*writer).Write([]byte("SHOW ME DA DETAILS"))

	// id := context.Param("id")
	// movie := queryMovie(&id)

	// if movie.IMDbId != nil {
	// 	context.IndentedJSON(http.StatusOK, &movie)
	// } else {
	// 	context.Status(http.StatusNotFound)
	// }
}
