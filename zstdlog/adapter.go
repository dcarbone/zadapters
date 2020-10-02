package zstdlog

import (
	stdlog "log"

	"github.com/rs/zerolog"
)

type adapter struct {
	level zerolog.Level
	def   bool
	log   zerolog.Logger
}

var (
	DefaultLevel = zerolog.NoLevel
)

// NewStdLoggerWithLevel will return an instance of *log.Logger where all messages will have the specified level
func NewStdLoggerWithLevel(logger zerolog.Logger, level zerolog.Level) *stdlog.Logger {
	return stdlog.New(adapter{level, false, logger}, "", 0)
}

// NewStdLogger will return an instance of *log.Logger where all messages will have no level attached
func NewStdLogger(logger zerolog.Logger) *stdlog.Logger {
	return stdlog.New(adapter{DefaultLevel, false, logger}, "", 0)
}

func (a adapter) Write(p []byte) (int, error) {
	if a.def {
		return a.WriteLevel(DefaultLevel, p)
	} else {
		return a.WriteLevel(a.level, p)
	}
}

func (a adapter) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	n := len(p)
	if n > 0 && p[n-1] == '\n' {
		p = p[0 : n-1]
	}
	a.log.WithLevel(level).Msg(string(p))
	return n, nil
}
