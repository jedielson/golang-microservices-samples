package ulog

import (
	"io/ioutil"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once     sync.Once
	instance *logrus.Logger
)

type OptionsFn func(*logrus.Logger)

func SetDiscardOutput() OptionsFn {
	return func(log *logrus.Logger) {
		log.Out = ioutil.Discard
	}
}

func New(fns ...OptionsFn) *logrus.Logger {
	log := logrus.New()
	for _, fn := range fns {
		fn(log)
	}
	return log
}

func Single(fns ...OptionsFn) *logrus.Logger {
	once.Do(func() {
		instance = New(fns...)
	})
	return instance
}
