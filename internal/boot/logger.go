package boot

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(
		zerolog.ConsoleWriter{Out: os.Stderr},
	) // human-friendly logging without efficiency
	log.Info().Msg("Logger initialized")
}
