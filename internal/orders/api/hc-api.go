package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

const apiPing = "/api/hc/ping"

func NewHcApi(r *mux.Router) {

	r.HandleFunc(apiPing, Ping()).Methods(http.MethodGet)
}

func Ping() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, "Pong")
	}
}
