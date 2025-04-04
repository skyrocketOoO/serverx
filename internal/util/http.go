package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/domain"
	"github.com/skyrocketOoO/serverx/internal/domain/err"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

func ParseValidate[T any](c *gin.Context, req *T) bool {
	if c.Request.Body == nil || c.Request.ContentLength == 0 {
		RespErr(c, http.StatusBadRequest, erx.W(err.New(err.EmptyRequest)))
		return false
	}

	if err := c.ShouldBindJSON(req); err != nil {
		RespErr(c, http.StatusBadRequest, erx.W(err))
		return false
	}

	if err := validate.Get().Struct(req); err != nil {
		RespErr(c, http.StatusBadRequest, erx.W(err))
		return false
	}
	return true
}

func RespErr(c *gin.Context, statusCode int, err error) {
	clientMsg, ok := erx.GetClientMsg(err)
	if !ok {
		c.JSON(statusCode, domain.ErrResp{Error: err.Error()})
		return
	}

	c.JSON(statusCode, domain.ErrResp{ID: clientMsg.ID, Error: clientMsg.Err})

	fullMsg, _ := erx.GetFullMsg(err)
	log.Error().Msg(fullMsg.ID)
	log.Error().Msg(fullMsg.Err)
	log.Error().Msg(fullMsg.CallStack)
}
