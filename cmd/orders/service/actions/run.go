package actions

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jedielson/jaeger-sample/internal/umux"
	"github.com/movidesk/go-gracefully"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	r := mux.NewRouter()
	r.Use(umux.PanicMiddleware)

	r.HandleFunc("/api/teste", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, "Olá Mundo")
	}).Methods("GET")

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				fmt.Printf("Falha na execução da aplicação\n")
			}
		}
	}()

	grace := gracefully.New(
		gracefully.WithTimeout(time.Second*5),
		gracefully.WithShutdown(srv),
	)

	return grace.Grace()
}
