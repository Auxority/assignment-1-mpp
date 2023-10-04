package router

import (
	"mpp/command"
	"net/http"
)

func DeleteMovie(writer http.ResponseWriter, request *http.Request) {
	id := GetUrlId(request)
	command.DeleteMovie(&id)
	writer.WriteHeader(http.StatusNoContent)
}
