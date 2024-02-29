package logging

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogging(logLevel string, humanReadable bool) {
	if humanReadable {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	loglvl, err := zerolog.ParseLevel(logLevel)

	if err != nil {
		//nolint:forbidigo // Allow `fmt` if issue arises before logger is fully configured
		fmt.Println("Failed to parse log level")
		os.Exit(1)
	}
	log.Info().Msgf("Setting loglevel to %s", loglvl)
	zerolog.SetGlobalLevel(loglvl)
}
