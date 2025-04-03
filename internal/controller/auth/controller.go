package auth

import (
	"context"

	"github.com/skyrocketOoO/serverx/internal/usecase/auth"
)

type Usecase interface {
	Login(c context.Context, in auth.LoginInput) (string, error)
	Register(c context.Context, in auth.RegisterInput) error
	ForgotPassword(c context.Context, in auth.ForgotPasswordInput) error
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
