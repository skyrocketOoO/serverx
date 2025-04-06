package boot

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/serverx/internal/domain"
	"github.com/skyrocketOoO/serverx/internal/service/loki"
)

func InitLogger() error {
	if domain.Env == "local" || domain.Env == "dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Info().Msg("Logger initialized")

	switch domain.LogTo {
	case "stdout":
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
			FormatTimestamp: func(i any) string {
				if i == nil {
					return "0000-00-00 00:00:00"
				}
				return i.(string)
			},
			FormatLevel: func(i any) string {
				if i == nil {
					return "[???]"
				}
				return "[" + i.(string) + "]"
			},
			FormatCaller: func(i any) string {
				if i == nil {
					return "unknown:0"
				}
				return simplifyCaller(i.(string))
			},
			FormatMessage: func(i any) string {
				if i == nil {
					return ""
				}
				return i.(string)
			},
			// NoColor: false,
		}
		log.Logger = zerolog.New(consoleWriter).With().Caller().Timestamp().Logger()
	case "loki":
		lokiWriter, err := loki.NewLokiWriter()
		if err != nil {
			return err
		}
		log.Logger = log.Output(lokiWriter).With().Caller().Timestamp().Logger()
	}

	return nil
}

func simplifyCaller(caller string) string {
	file := filepath.Base(caller)
	dir := filepath.Dir(caller)

	return filepath.Join(filepath.Base(dir), file)
}
