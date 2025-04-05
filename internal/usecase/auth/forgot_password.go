package authucase

import "context"

type ForgotPasswordIn struct {
	Email string `json:"email" validate:"required"`
}

func (u *Usecase) ForgotPassword(c context.Context, in ForgotPasswordIn) error {
	return nil
}
