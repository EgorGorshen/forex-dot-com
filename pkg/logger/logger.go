package logger

import (
	"os"
	"time"

	"github.com/EgorGorshen/forex-dot-com/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339

	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).
		Level(config.Config.LogLevel).
		With().
		Timestamp().
		Caller().
		Logger()
}
