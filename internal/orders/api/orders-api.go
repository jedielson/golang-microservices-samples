package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type OrdersApiUrls struct {
	Place string
}

func (s OrdersApiUrls) Get() OrdersApiUrls {
	return OrdersApiUrls{
		Place: "/api/orders/place",
	}
}

func NewOrdersApi(r *mux.Router) {

	routes := OrdersApiUrls{}.Get()
	r.HandleFunc(routes.Place, Place()).Methods("POST")
}

func Place() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, "Olá Mundo")
	}
}
