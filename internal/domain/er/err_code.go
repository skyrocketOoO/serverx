package er

import "net/http"

// Use string as code to support more flexibility than float
type code string

func (c code) String() string {
	return string(c)
}

const (
	BadRequest    code = "400.0000"
	EmptyRequest  code = "400.0001"
	ParsePayload  code = "400.0002"
	ValidateInput code = "400.0003"

	Unauthorized        code = "401.0000"
	NewPasswordRequired code = "401.0001"

	NotFound code = "404.0000"

	Unknown       code = "500.0000"
	DBUnavailable code = "500.0001"

	NotImplemented code = "501.0000"
)

var CodeToMsg = map[code]string{
	BadRequest:    "bad request",
	EmptyRequest:  "empty request body",
	ParsePayload:  "parse payload error",
	ValidateInput: "validate input error",

	Unauthorized:        "unauthorized",
	NewPasswordRequired: "new password required",

	NotFound: "not found",

	Unknown:       "unknown",
	DBUnavailable: "database unavailable",

	NotImplemented: "not implemented",
}

var codeToHTTP = map[code]int{
	BadRequest:    http.StatusBadRequest,
	EmptyRequest:  http.StatusBadRequest,
	ParsePayload:  http.StatusBadRequest,
	ValidateInput: http.StatusBadRequest,

	Unauthorized:        http.StatusUnauthorized,
	NewPasswordRequired: http.StatusUnauthorized,

	NotFound: http.StatusNotFound,

	Unknown:       http.StatusInternalServerError,
	DBUnavailable: http.StatusInternalServerError,

	NotImplemented: http.StatusNotImplemented,
}
