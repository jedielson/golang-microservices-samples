package actions

import (
	"time"

	"github.com/jedielson/jaeger-sample/internal/common"
	"github.com/jedielson/jaeger-sample/internal/orders/api"
	"github.com/jedielson/jaeger-sample/internal/umux"
	"github.com/movidesk/go-gracefully"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {

	locator := common.Newlocator()
	locator.Register(c)
	locator.RegisterLogrus(c)

	r := umux.NewMuxRouter()
	api.NewOrdersApi(r)
	api.NewHcApi(r)

	srv := umux.GetServer(locator, r)

	log := locator.FindLogrus()
	log.
		WithField("app", locator.FindApplication()).
		WithField("env", locator.FindEnv()).
		WithField("addr", locator.FindAddr()).
		WithField("ver", locator.FindVersion()).
		Info("Listening")

	go umux.Serve(srv, r)

	grace := gracefully.New(
		gracefully.WithTimeout(time.Second*5),
		gracefully.WithShutdown(srv),
	)

	return grace.Grace()
}
