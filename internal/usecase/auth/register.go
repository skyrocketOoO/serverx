package authucase

import "context"

type RegisterInput struct {
	Email    string `validate:"required"`
	NickName string `validate:"required"`
}

func (u *Usecase) Register(c context.Context, in RegisterInput) error {
	return nil
}
