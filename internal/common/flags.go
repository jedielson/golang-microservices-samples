package common

import "github.com/urfave/cli/v2"

var (
	AmqpDsnFlag = &cli.StringFlag{
		Name:     "amqp-dsn",
		Usage:    "amqp dsn to use for connecting on the amqp cluster",
		Value:    "amqp://guest:guest@localhost:5672",
		EnvVars:  []string{"MS_AMQP_DSN"},
		Required: false,
	}
	LocalFlag = &cli.BoolFlag{
		Name:     "movidesk-local",
		Usage:    "Check if is running in local environment (using localhost)",
		Value:    false,
		EnvVars:  []string{"MS_LOCAL"},
		Required: false,
	}
	PortFlag = &cli.StringFlag{
		Name:     "port",
		Usage:    "port to bind for serving the application",
		EnvVars:  []string{"MS_PORT"},
		Required: true,
	}
)
