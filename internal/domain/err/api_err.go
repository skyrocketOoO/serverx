package err

type APIError struct {
	TraceID string  `json:"traceID"`
	Code    float32 `json:"code"`
	Msg     string  `json:"msg"`
}

func NewAPIError(traceID string, code float32, msg string) *APIError {
	return &APIError{
		TraceID: traceID,
		Code:    code,
		Msg:     msg,
	}
}
