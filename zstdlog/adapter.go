package zstdlog

import (
	"github.com/rs/zerolog"
	"strings"
)

type (
	// Adapter is a simple wrapper for a ZeroLog logger that allows it to be passed to things which
	// expect an instances of log.Logger
	Adapter struct {
		l  zerolog.Logger
		ev *zerolog.Event
	}
)

// New constructs a new Adapter with the provided zerolog.Logger as the writer
func New(logger zerolog.Logger) *Adapter {
	return &Adapter{l: logger}
}

// NewWithLevel will log using the specified level
func NewWithLevel(logger zerolog.Logger, level zerolog.Level) *Adapter {
	var ev *zerolog.Event
	switch level {
	case zerolog.DebugLevel:
		ev = logger.Debug()
	case zerolog.InfoLevel:
		ev = logger.Info()
	case zerolog.WarnLevel:
		ev = logger.Warn()
	case zerolog.ErrorLevel:
		ev = logger.Error()
	case zerolog.FatalLevel:
		ev = logger.Fatal()
	case zerolog.PanicLevel:
		ev = logger.Panic()
	}
	return &Adapter{l: logger, ev: ev}
}

func (a *Adapter) Print(v ...interface{}) {
	if a.ev == nil {
		a.l.Print(v...)
	} else {
		a.Printf(strings.TrimRight(strings.Repeat("%v ", len(v)), " "), v...)
	}
}

func (a *Adapter) Println(v ...interface{}) {
	if a.ev == nil {
		a.l.Print(v...)
	} else {
		a.Printf(strings.TrimRight(strings.Repeat("%v ", len(v)), " ")+"\n", v...)
	}
}

func (a *Adapter) Printf(f string, v ...interface{}) {
	if a.ev == nil {
		a.l.Printf(f, v...)
	} else {
		a.ev.Msgf(f, v...)
	}
}
