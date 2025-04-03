package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/global"
)

func ParseValidate[T any](c *gin.Context, req *T) bool {
	if c.Request.Body == nil || c.Request.ContentLength == 0 {
		RespErr(c, http.StatusBadRequest, erx.W(ErrEmptyRequest))
		return false
	}

	if err := c.ShouldBindJSON(req); err != nil {
		RespErr(c, http.StatusBadRequest, erx.W(err))
		return false
	}

	if err := global.Validator.Struct(req); err != nil {
		RespErr(c, http.StatusBadRequest, erx.W(err))
		return false
	}
	return true
}

type ErrResp struct {
	ID    string `json:"id"`
	Error string `json:"error"`
}

func RespErr(c *gin.Context, statusCode int, err error) {
	clientMsg, ok := erx.GetClientMsg(err)
	if !ok {
		c.JSON(statusCode, ErrResp{Error: err.Error()})
		return
	}

	c.JSON(statusCode, ErrResp{ID: clientMsg.ID, Error: clientMsg.Err})

	fullMsg, _ := erx.GetFullMsg(err)
	log.Error().Msg(fullMsg.ID)
	log.Error().Msg(fullMsg.Err)
	log.Error().Msg(fullMsg.CallStack)
}
