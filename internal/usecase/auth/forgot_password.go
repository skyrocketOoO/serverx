package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type ForgotPasswordIn struct {
	Email string `json:"email" validate:"required"`
}

func (u *Usecase) ForgotPassword(c context.Context, in ForgotPasswordIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.ForgotPasswordInput{
		ClientId:   aws.String(u.cognitoSvc.ClientID),
		SecretHash: aws.String(u.cognitoSvc.ComputeSecretHash(in.Email)),
		Username:   aws.String(in.Email),
	}

	if _, err := u.cognitoSvc.Client.ForgotPassword(c, input); err != nil {
		return er.W(err)
	}

	return nil
}
