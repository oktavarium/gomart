package log

import (
	"os"

	"github.com/rs/zerolog"
)

type Log struct {
	logger zerolog.Logger
}

func NewLogger(logLevel string) *Log {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(level)

	return &Log{logger: zerolog.New(os.Stdout)}
}

func (l *Log) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *Log) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Log) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *Log) Error(err error) {
	l.logger.Debug().Msg(err.Error())
}

func (l *Log) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}
