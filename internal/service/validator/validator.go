package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

var validate *validator.Validate

func Get() *validator.Validate {
	return validate
}

func New() {
	log.Info().Msg("InitValidator")
	validate = validator.New(validator.WithRequiredStructEnabled())
}
