package zstdlog

import (
	stdlog "log"

	"github.com/rs/zerolog"
)

type adapter struct {
	ev *zerolog.Event
}

// NewStdLoggerWithLevel will return an instance of *log.Logger where all messages will have the specified level
func NewStdLoggerWithLevel(logger zerolog.Logger, level zerolog.Level) *stdlog.Logger {
	return stdlog.New(adapter{logger.WithLevel(level)}, "", 0)
}

// NewStdLogger will return an instance of *log.Logger where all messages will have no level attached
func NewStdLogger(logger zerolog.Logger) *stdlog.Logger {
	return NewStdLoggerWithLevel(logger, zerolog.NoLevel)
}

func (a adapter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		p = p[0 : n-1]
	}
	a.ev.Msg(string(p))
	return
}

// Deprecated
type Adapter struct {
	*stdlog.Logger
}

// Deprecated
func New(logger zerolog.Logger) *Adapter {
	return &Adapter{NewStdLogger(logger)}
}

// Deprecated
func NewWithLevel(logger zerolog.Logger, level zerolog.Level) *Adapter {
	return &Adapter{NewStdLoggerWithLevel(logger, level)}
}

// Deprecated
func (a *Adapter) Print(v ...interface{}) { a.Logger.Print(v...) }

// Deprecated
func (a *Adapter) Println(v ...interface{}) { a.Logger.Println(v...) }

// Deprecated
func (a *Adapter) Printf(f string, v ...interface{}) { a.Logger.Printf(f, v...) }
