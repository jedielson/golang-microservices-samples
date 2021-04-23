package main

import (
	"log"
	"os"

	"github.com/jedielson/jaeger-sample/cmd/orders/service/actions"
	"github.com/urfave/cli/v2"
)

var (
	AppName    = "orders-service"
	AppUsage   = ""
	AppVersion = "0.0.1"
	GitSummary = "none"
	GitBranch  = "none"
	GitMerge   = "0"
	CiBuild    = "0"
)

func main() {

	cli.VersionPrinter = func(c *cli.Context) {
		log.Printf("version=%s summary=%s branch=%s merge=%s build=%s", c.App.Version, GitSummary, GitBranch, GitMerge, CiBuild)
	}

	app := &cli.App{
		Name:    AppName,
		Usage:   AppUsage,
		Version: AppVersion,
		Action:  actions.Run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}
