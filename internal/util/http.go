package util

import (
	"github.com/gin-gonic/gin"
	er "github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

func ParseValidate[T any](c *gin.Context, req *T) bool {
	if c.Request.Body == nil || c.Request.ContentLength == 0 {
		er.Bind(c, er.NewAppErr(er.EmptyRequest))
		return false
	}

	if err := c.ShouldBindJSON(req); err != nil {
		er.Bind(c, er.W(err, er.ParsePayload))
		return false
	}

	if err := validate.Get().Struct(req); err != nil {
		er.Bind(c, er.W(err, er.ValidateInput))
		return false
	}
	return true
}
