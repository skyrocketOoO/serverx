package auth

import "context"

type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required"`
}

func (u *Usecase) ForgotPassword(c context.Context, in ForgotPasswordInput) error {
	return nil
}
