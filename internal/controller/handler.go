package controller

import (
	"github.com/skyrocketOoO/serverx/internal/controller/auth"
	"github.com/skyrocketOoO/serverx/internal/controller/general"
	_ "github.com/skyrocketOoO/serverx/internal/controller/middleware" // avoid import cycle
)

type Handler struct {
	Auth    auth.Handler
	General general.Handler
}

func NewHandler(
	auth auth.Handler,
	general general.Handler,
) *Handler {
	return &Handler{
		Auth:    auth,
		General: general,
	}
}
