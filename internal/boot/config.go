package boot

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/serverx/internal/domain"
)

func InitConfig() (err error) {
	if domain.Env == "local" {
		if err := godotenv.Load(".env"); err != nil {
			log.Error().Msg("Error loading .env file")
			return err
		}
	}

	return nil
}
