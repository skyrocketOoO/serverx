package controller

import (
	_ "github.com/skyrocketOoO/serverx/internal/controller/middleware" // avoid import cycle
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
