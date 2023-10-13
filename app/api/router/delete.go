package router

import (
	"mpp/api/command"
	"mpp/error_util"
	"net/http"
)

func DeleteMovie(writer http.ResponseWriter, request *http.Request) {
	id := GetUrlId(request)
	err := command.DeleteMovie(&id)
	error_util.CheckError(err)
	writer.WriteHeader(http.StatusNoContent)
}
