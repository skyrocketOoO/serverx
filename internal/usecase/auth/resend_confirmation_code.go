package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type ResendConfirmationCodeIn struct {
	Email string `json:"email" validate:"required,email"`
}

func (u *Usecase) ResendConfirmationCode(c context.Context, in ResendConfirmationCodeIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId:   aws.String(u.cognitoSvc.ClientID),
		Username:   aws.String(in.Email),
		SecretHash: aws.String(u.cognitoSvc.ComputeSecretHash(in.Email)),
	}

	_, err := u.cognitoSvc.Client.ResendConfirmationCode(c, input)
	if err != nil {
		return er.W(err, er.BadRequest)
	}

	return nil
}
