package main

import (
	"github.com/EgorGorshen/forex-dot-com/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.InitLogger()

	log.Info().Msg("Hello, Egor")
}
