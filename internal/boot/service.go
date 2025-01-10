package boot

import (
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/web-server-template/internal/service/db"
	"github.com/skyrocketOoO/web-server-template/internal/service/validator"
)

func NewService() error {
	log.Info().Msg("InitService")
	if err := db.New("sqlite"); err != nil {
		return err
	}

	validator.New()

	return nil
}