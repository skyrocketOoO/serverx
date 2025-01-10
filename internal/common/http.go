package cm

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	dm "github.com/skyrocketOoO/serverx/internal/global/domain"
	"github.com/skyrocketOoO/serverx/internal/service/inter/validator"
)

func BindAndValidate[T any](c *gin.Context, req *T) bool {
	if c.Request.Body == nil || c.Request.ContentLength == 0 {
		dm.RespErr(c, http.StatusBadRequest, erx.W(dm.ErrEmptyRequest))
		return false
	}

	if err := c.ShouldBindJSON(req); err != nil {
		dm.RespErr(c, http.StatusBadRequest, erx.W(err))
		return false
	}

	if err := validator.Get().Struct(req); err != nil {
		dm.RespErr(c, http.StatusBadRequest, erx.W(err))
		return false
	}
	return true
}

func RespErr(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, dm.ErrResp{Error: err.Error()})
}
