package domain

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

const (
	ApiVersion = "v1"
)

var (
	// env | flag
	Database    string
	AutoMigrate bool = false
	Env         string

	Validator *validator.Validate // use a single instance of Validate, it caches struct info
)

var (
	ErrNotImplement     = errors.New("not implemented")
	ErrEmptyRequest     = errors.New("empty request body")
	ErrUnknown          = errors.New("unknown")
	ErrLoginFailed      = errors.New("login failed")
	ErrUserNameRepetite = errors.New("user name repetite")
)

type ErrResp struct {
	ID    string `json:"id"`
	Error string `json:"error"`
}
