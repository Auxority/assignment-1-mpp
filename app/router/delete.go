package router

import "net/http"

func OnMovieDeleteRequest(writer *http.ResponseWriter, request *http.Request) {
	(*writer).Write([]byte("DELETE DA MOVIE"))
	// id := context.Param("id")
	// command.DeleteMovie(&id)
	// context.Status(http.StatusNoContent)
}
