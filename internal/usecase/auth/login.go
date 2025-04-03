package auth

import "context"

type LoginInput struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func (u *Usecase) Login(c context.Context, in LoginInput) (string, error) {
	return "", nil
}
