package zgocb

import (
	"github.com/couchbase/gocb"
	"github.com/rs/zerolog"
)

type (
	// LevelMap is used to translate GOCB log levels to ZeroLog levels
	LevelMap map[gocb.LogLevel]zerolog.Level

	// Adapter implements gocb.Logger, allowing for a ZeroLog logger to be used.
	Adapter struct {
		l  zerolog.Logger
		lm LevelMap
	}
)

var (
	// This is the default map that will be used if Adapter is constructed with "NewDefault"
	// Change as makes sense to your app.
	DefaultLevelMap = LevelMap{
		gocb.LogMaxVerbosity: zerolog.Disabled, // this is VERY chatty, only enable if you really mean it
		gocb.LogSched:        zerolog.Disabled,
		gocb.LogTrace:        zerolog.Disabled,
		gocb.LogDebug:        zerolog.DebugLevel,
		gocb.LogInfo:         zerolog.InfoLevel,
		gocb.LogWarn:         zerolog.WarnLevel,
		gocb.LogError:        zerolog.ErrorLevel,
	}
)

// New initializes a new Adapter with the specified level map
func New(levelMap LevelMap, logger zerolog.Logger) *Adapter {
	// just in case...
	if levelMap == nil {
		levelMap = DefaultLevelMap
	}
	return &Adapter{
		l:  logger,
		lm: levelMap,
	}
}

// NewDefault creates a new adapter with the default LevelMap
func NewDefault(logger zerolog.Logger) *Adapter {
	return New(DefaultLevelMap, logger)
}

// Log translates the gocb log level to a zerolog event based upon the event map created with the Adapter
func (a *Adapter) Log(level gocb.LogLevel, offset int, f string, v ...interface{}) error {
	if l, ok := a.lm[level]; ok {
		switch l {
		case zerolog.DebugLevel:
			a.l.Debug().Msgf(f, v...)
		case zerolog.InfoLevel:
			a.l.Info().Msgf(f, v...)
		case zerolog.WarnLevel:
			a.l.Warn().Msgf(f, v...)
		case zerolog.ErrorLevel:
			a.l.Error().Msgf(f, v...)
		case zerolog.FatalLevel:
			a.l.Fatal().Msgf(f, v...)
		case zerolog.PanicLevel:
			a.l.Panic().Msgf(f, v...)
		case zerolog.NoLevel:
			a.l.Printf(f, v...)
		case zerolog.Disabled:
			// do nothing
		}
	}
	return nil
}