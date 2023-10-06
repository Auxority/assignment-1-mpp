package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetUrlId(request *http.Request) string {
	return GetUrlParameter(request, "id")
}

func GetUrlParameter(request *http.Request, name string) string {
	return mux.Vars(request)[name]
}
