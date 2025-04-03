package domain

import (
	"errors"
)

const (
	ApiVersion = "v1"
)

var (
	// env | flag
	Database    string
	AutoMigrate bool = false
	Env         string
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
