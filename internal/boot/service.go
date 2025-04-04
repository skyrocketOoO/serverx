package boot

import (
	"github.com/rs/zerolog/log"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

func NewService() error {
	log.Info().Msg("InitService")
	// if err := postgres.New(); err != nil {
	// 	return err
	// }

	validate.New()

	return nil
}
