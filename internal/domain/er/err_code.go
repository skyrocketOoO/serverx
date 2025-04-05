package er

import "net/http"

// Use string as code to support more flexibility than float
type code string

func (c code) String() string {
	return string(c)
}

const (
	EmptyRequest  code = "400.0001"
	ParsePayload  code = "400.0002"
	ValidateInput code = "400.0003"

	NewPasswordRequired code = "401.0001"

	NotFound code = "404"

	Unknown       code = "500"
	DBUnavailable code = "500.0001"

	NotImplemented code = "501"
)

var CodeToMsg = map[code]string{
	EmptyRequest:  "empty request body",
	ParsePayload:  "parse payload error",
	ValidateInput: "validate input error",

	NewPasswordRequired: "new password required",

	NotFound: "not found",

	Unknown:       "unknown",
	DBUnavailable: "database unavailable",

	NotImplemented: "not implemented",
}

var codeToHTTP = map[code]int{
	EmptyRequest:  http.StatusBadRequest,
	ParsePayload:  http.StatusBadRequest,
	ValidateInput: http.StatusBadRequest,

	NewPasswordRequired: http.StatusUnauthorized,
	NotFound:            http.StatusNotFound,

	Unknown:       http.StatusInternalServerError,
	DBUnavailable: http.StatusInternalServerError,

	NotImplemented: http.StatusNotImplemented,
}
