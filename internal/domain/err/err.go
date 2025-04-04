package err

import "net/http"

type CustomErr struct {
	code Code
	msg  string
}

func (e *CustomErr) Error() string {
	return e.msg
}

func (e *CustomErr) Code() Code {
	return e.code
}

func (e *CustomErr) HTTPCode() int {
	switch e.code {
	case EmptyRequest:
		return http.StatusBadRequest
	case NewPasswordRequired:
		return http.StatusUnauthorized
	case NotImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}
}

func New(code Code) error {
	return &CustomErr{code: code, msg: CodeToMsg[code]}
}
