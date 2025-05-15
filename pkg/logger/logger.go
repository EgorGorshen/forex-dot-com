package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitLogger() {
	viper.AutomaticEnv()

	levelStr := viper.GetString("LOG_LEVEL")
	level, err := zerolog.ParseLevel(strings.ToLower(levelStr))
	if err != nil {
		level = zerolog.InfoLevel
	}

	zerolog.TimeFieldFormat = time.RFC3339

	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}).
		Level(level).
		With().
		Timestamp().
		Caller().
		Logger()
}
