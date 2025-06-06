package authcontroller

import (
	"context"

	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
)

type Usecase interface {
	Login(c context.Context, in authucase.LoginIn) (*authucase.LoginOut, error)
	SignUp(c context.Context, in authucase.SignUpIn) error
	ConfirmSignUp(c context.Context, in authucase.ConfirmSignUpIn) error
	ForgotPassword(c context.Context, in authucase.ForgotPasswordIn) error
	ConfirmForgotPassword(c context.Context, in authucase.ConfirmForgotPasswordIn) error
	RefreshToken(c context.Context, in authucase.RefreshTokenIn) (*authucase.RefreshTokenOut, error)
	ChangePassword(c context.Context, in authucase.ChangePasswordIn) error
	ResendConfirmationCode(c context.Context, in authucase.ResendConfirmationCodeIn) error
	InviteUser(c context.Context, in authucase.InviteUserIn) error
	SetNewPassword(c context.Context, in authucase.SetNewPasswordIn) error
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
