package authcontroller

import (
	"context"

	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
)

type Usecase interface {
	Login(c context.Context, in authucase.LoginIn) (*authucase.LoginOut, error)
	SignUp(c context.Context, in authucase.SignUpInput) error
	ForgotPassword(c context.Context, in authucase.ForgotPasswordInput) error
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
