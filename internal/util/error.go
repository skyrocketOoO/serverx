package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/domain"
	"gorm.io/gorm"
)

// Convert original error to http status code
func ToHttpCode(err error) int {
	ogErr := erx.RUnwrap(err)
	switch ogErr {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case domain.ErrNotImplement:
		return http.StatusNotImplemented
	case domain.ErrEmptyRequest:
		return http.StatusBadRequest
	case domain.ErrUnknown:
		return http.StatusInternalServerError
	case domain.ErrLoginFailed:
		return http.StatusUnauthorized
	case domain.ErrUserNameRepetite:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func RespErrWithCode(c *gin.Context, err error) {
	statusCode := ToHttpCode(err)
	RespErr(c, statusCode, err)
}
