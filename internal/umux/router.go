package umux

import "github.com/gorilla/mux"

func NewMuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(PanicMiddleware)

	return r
}
