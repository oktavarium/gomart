package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(logLevel string) {
	l, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		l = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(l)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Error(err error) {
	log.Debug().Msg(err.Error())
}
