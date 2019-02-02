package zstdlog

import (
	"strings"

	"github.com/rs/zerolog"
)

type (
	// Adapter is a simple wrapper for a ZeroLog logger that allows it to be passed to things which
	// expect an instances of log.Logger
	Adapter struct {
		l  zerolog.Logger
		ev *zerolog.Event
	}
)

// NewWithLevel will log using the specified level
func NewWithLevel(logger zerolog.Logger, level zerolog.Level) *Adapter {
	return &Adapter{l: logger.Level(level), ev: logger.WithLevel(level)}
}

// New constructs a new Adapter with the provided zerolog.Logger as the writer
func New(logger zerolog.Logger) *Adapter {
	return NewWithLevel(logger, zerolog.NoLevel)
}

func (a *Adapter) Print(v ...interface{}) {
	a.ev.Msgf(strings.TrimRight(strings.Repeat("%v ", len(v)), " "), v...)
}

func (a *Adapter) Println(v ...interface{}) {
	a.ev.Msgf(strings.TrimRight(strings.Repeat("%v ", len(v)), " ")+"\n", v...)
}

func (a *Adapter) Printf(f string, v ...interface{}) {
	a.ev.Msgf(f, v...)
}
