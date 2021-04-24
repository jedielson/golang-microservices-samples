package common

import (
	"github.com/jedielson/jaeger-sample/internal/ulog"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type Locator interface {
	Register(*cli.Context)
	FindEnv() string
	FindVersion() string
	FindApplication() string
	FindAddr() string

	RegisterLogrus(*cli.Context)
	FindLogrus() *logrus.Logger
}

type locator struct {
	env         string
	version     string
	application string
	addr        string

	logrus *logrus.Logger
}

func Newlocator() Locator {
	return &locator{}
}

func (l *locator) Register(ctx *cli.Context) {
	isLocal := ctx.Bool(LocalFlag.Name)
	l.env = "production"
	if isLocal {
		l.env = "development"
	}

	l.addr = ctx.String(PortFlag.Name)
	l.version = ctx.App.Version
	l.application = ctx.App.Name
}

func (l *locator) FindEnv() string {
	if len(l.env) == 0 {
		panic("env was not registered")
	}
	return l.env
}

func (l *locator) FindVersion() string {
	if len(l.version) == 0 {
		panic("version was not registered")
	}
	return l.version
}

func (l *locator) FindApplication() string {
	if len(l.application) == 0 {
		panic("application was not registered")
	}
	return l.application
}

func (l *locator) FindAddr() string {
	if len(l.addr) == 0 {
		panic("addr was not registered")
	}
	return l.addr
}

func (l *locator) RegisterLogrus(ctx *cli.Context) {
	if l.logrus != nil {
		return
	}
	l.logrus = ulog.Single()
}

func (l *locator) FindLogrus() *logrus.Logger {
	if l.logrus == nil {
		panic("logrus was not registered")
	}
	return l.logrus
}
