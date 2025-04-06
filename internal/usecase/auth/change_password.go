package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type ChangePasswordIn struct {
	AccessToken string `json:"accessToken" validate:"required"`
	OldPass     string `json:"oldPass"     validate:"required"`
	NewPass     string `json:"newPass"     validate:"required"`
}

func (u *Usecase) ChangePassword(c context.Context, in ChangePasswordIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      aws.String(in.AccessToken),
		PreviousPassword: aws.String(in.OldPass),
		ProposedPassword: aws.String(in.NewPass),
	}

	if _, err := u.cognitoSvc.Client.ChangePassword(c, input); err != nil {
		return er.W(err, er.BadRequest)
	}

	return nil
}
