package authcontroller

import (
	"context"

	authusecase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
)

type Usecase interface {
	Login(c context.Context, in authusecase.LoginInput) (string, error)
	Register(c context.Context, in authusecase.RegisterInput) error
	ForgotPassword(c context.Context, in authusecase.ForgotPasswordInput) error
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
