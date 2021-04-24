package actions

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jedielson/jaeger-sample/internal/common"
	"github.com/jedielson/jaeger-sample/internal/umux"
	"github.com/movidesk/go-gracefully"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {

	locator := common.Newlocator()
	locator.Register(c)
	locator.RegisterLogrus(c)

	r := umux.NewMuxRouter()

	r.HandleFunc("/api/teste", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, "Olá Mundo")
	}).Methods("GET")

	log := locator.FindLogrus()
	log.
		WithField("app", locator.FindApplication()).
		WithField("env", locator.FindEnv()).
		WithField("addr", locator.FindAddr()).
		WithField("ver", locator.FindVersion()).
		Info("Listening")

	srv := &http.Server{
		Addr:         locator.FindAddr(),
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
