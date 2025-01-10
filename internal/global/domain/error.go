package dm

import "errors"

var (
	ErrEmptyRequest = errors.New("empty request body")
	ErrUnknown      = errors.New("unknown")
)
