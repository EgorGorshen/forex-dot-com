package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Conf struct {
	LogLevel zerolog.Level
}

var Config = newConfig()

func newConfig() Conf {
	return Conf{
		LogLevel: zerolog.TraceLevel,
	}
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		panic(".env файл не найден, пропускаем")
	}
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "DEBUG")

	Config.setLogLevel()
}

// Set LogLevel
func (c *Conf) setLogLevel() {
	levelStr := viper.GetString("LOG_LEVEL")
	level, err := zerolog.ParseLevel(strings.ToLower(levelStr))
	if err != nil {
		level = zerolog.TraceLevel
	}
	c.LogLevel = level
}
