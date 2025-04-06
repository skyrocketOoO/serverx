package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type ConfirmSignUpIn struct {
	Email string `json:"email" validate:"required"`
	Code  string `json:"code"  validate:"required"`
}

func (u *Usecase) ConfirmSignUp(c context.Context, in ConfirmSignUpIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(u.cognitoSvc.ClientID),
		Username:         aws.String(in.Email),
		ConfirmationCode: aws.String(in.Code),
		SecretHash:       aws.String(u.cognitoSvc.ComputeSecretHash(in.Email)),
	}

	if _, err := u.cognitoSvc.Client.ConfirmSignUp(c, input); err != nil {
		return er.W(err)
	}

	return nil
}
