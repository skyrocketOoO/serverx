package authucase

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type SignUpIn struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
	NickName string `validate:"required"`
}

func (u *Usecase) SignUp(c context.Context, in SignUpIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(u.cognitoSvc.ClientID),
		Username: aws.String(in.Email),
		Password: aws.String(in.Password),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(in.Email)},
			{Name: aws.String("nickname"), Value: aws.String(in.NickName)},
		},
		SecretHash: aws.String(u.cognitoSvc.ComputeSecretHash(in.Email)),
	}

	if _, err := u.cognitoSvc.Client.SignUp(c, input); err != nil {
		var e *types.InvalidPasswordException
		if errors.As(err, &e) {
			return er.W(err, er.BadRequest)
		}

		// log.Debug().Str("errType", fmt.Sprintf("%T", err)).Str("msg", err.Error()).Send()
		return er.W(err, er.BadRequest)
	}

	return nil
}
