package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func NewOrdersApi(r *mux.Router) {

	s := r.PathPrefix("/api/orders").Subrouter()
	s.HandleFunc("/place", Place()).Methods("POST")
}

func Place() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, "Olá Mundo")
	}
}
