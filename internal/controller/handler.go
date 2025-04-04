package controller

import (
	authcontroller "github.com/skyrocketOoO/serverx/internal/controller/auth"
	generalcontroller "github.com/skyrocketOoO/serverx/internal/controller/general"
	_ "github.com/skyrocketOoO/serverx/internal/controller/middleware" // avoid import cycle
)

type Handler struct {
	Auth    *authcontroller.Handler
	General *generalcontroller.Handler
}

func NewHandler(
	auth *authcontroller.Handler,
	general *generalcontroller.Handler,
) *Handler {
	return &Handler{
		Auth:    auth,
		General: general,
	}
}
