package util

// Convert original error to http status code
// func ToHttpCode(e error) int {
// 	ogErr := erx.RUnwrap(e)
// 	switch ogErr {
// 	case gorm.ErrRecordNotFound:
// 		return http.StatusNotFound
// 	case err.NotImplement:
// 		return http.StatusNotImplemented
// 	case domain.ErrEmptyRequest:
// 		return http.StatusBadRequest
// 	case domain.ErrUnknown:
// 		return http.StatusInternalServerError
// 	case domain.ErrLoginFailed:
// 		return http.StatusUnauthorized
// 	case domain.ErrUserNameRepetite:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }

// func RespErrWithCode(c *gin.Context, err error) {
// 	statusCode := ToHttpCode(err)
// 	RespErr(c, statusCode, err)
// }
