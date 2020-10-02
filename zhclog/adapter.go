package zhclog

import (
	"github.com/hashicorp/go-hclog"
	"github.com/rs/zerolog"
)

type (
	// LevelMap is used to translate hclog levels to zerolog levels
	LevelMap map[hclog.Level]zerolog.Level

	// Adapter implements hclog.SinkAdapter
	Adapter struct {
		l  zerolog.Logger
		lm LevelMap
	}
)

var (
	// DefaultLevelMap is used by any new Adapter that is not provided an explicit map.
	DefaultLevelMap = LevelMap{
		hclog.NoLevel: zerolog.NoLevel,
		hclog.Trace:   zerolog.TraceLevel,
		hclog.Debug:   zerolog.DebugLevel,
		hclog.Info:    zerolog.InfoLevel,
		hclog.Warn:    zerolog.WarnLevel,
		hclog.Error:   zerolog.ErrorLevel,
	}
)

func New(levelMap LevelMap, logger zerolog.Logger) *Adapter {
	if levelMap == nil {
		levelMap = DefaultLevelMap
	}
	return &Adapter{
		l:  logger,
		lm: levelMap,
	}
}

func (a *Adapter) Accept(name string, level hclog.Level, msg string, args ...interface{}) {
	var (
		l  zerolog.Level
		ok bool
	)
	if l, ok = a.lm[level]; !ok {
		return
	}

	switch l {
	case zerolog.TraceLevel:
		a.l.Trace().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.DebugLevel:
		a.l.Debug().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.InfoLevel:
		a.l.Info().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.WarnLevel:
		a.l.Warn().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.ErrorLevel:
		a.l.Error().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.FatalLevel:
		a.l.Fatal().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.PanicLevel:
		a.l.Panic().Str("hclog_name", name).Msgf(msg, args...)
	case zerolog.NoLevel:
		a.l.Printf(msg, args...)
	case zerolog.Disabled:
		// do nothing
	}
}
