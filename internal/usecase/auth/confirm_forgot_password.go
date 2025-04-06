package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type ConfirmForgotPasswordIn struct {
	Email   string `json:"email"   validate:"required"`
	Code    string `json:"code"    validate:"required"`
	NewPass string `json:"newPass" validate:"required"`
}

func (u *Usecase) ConfirmForgotPassword(
	c context.Context,
	in ConfirmForgotPasswordIn,
) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(u.cognitoSvc.ClientID),
		SecretHash:       aws.String(u.cognitoSvc.ComputeSecretHash(in.Email)),
		Username:         aws.String(in.Email),
		ConfirmationCode: aws.String(in.Code),
		Password:         aws.String(in.NewPass),
	}

	if _, err := u.cognitoSvc.Client.ConfirmForgotPassword(c, input); err != nil {
		return er.W(err)
	}

	return nil
}
