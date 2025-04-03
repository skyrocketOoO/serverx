package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/serverx/internal/domain"
)

func New() {
	log.Info().Msg("InitValidator")
	domain.Validator = validator.New(validator.WithRequiredStructEnabled())
}
