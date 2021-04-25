package umux

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jedielson/jaeger-sample/internal/common"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return handlers.ContentTypeHandler(next, "application/json")
}

func NewMuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(contentTypeMiddleware)
	r.Use(handlers.RecoveryHandler())
	r.Use(handlers.CompressHandler)

	return r
}

func GetServer(locator common.Locator, r *mux.Router) *http.Server {
	return &http.Server{
		Addr:         locator.FindAddr(),
		Handler:      r,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

func Serve(srv *http.Server, r *mux.Router) {
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("Falha na execução da aplicação\n")
		}
	}
}
