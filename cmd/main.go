package main

import (
	"github.com/EgorGorshen/forex-dot-com/config"
	"github.com/EgorGorshen/forex-dot-com/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	config.InitConfig()
	logger.InitLogger()

	log.Info().Msg("Hello")
}
