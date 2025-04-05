package authucase

import "context"

type ConfirmForgotPasswordIn struct {
	Email   string `json:"email"   validate:"required"`
	Code    string `json:"code"    validate:"required"`
	NewPass string `json:"newPass" validate:"required"`
}

func (u *Usecase) ConfirmForgotPassword(
	c context.Context,
	in ConfirmForgotPasswordIn,
) error {
	return nil
}
