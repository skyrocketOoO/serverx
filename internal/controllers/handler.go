package controller

import (
	_ "github.com/skyrocketOoO/serverx/internal/controllers/middlewares" // avoid import cycle
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
