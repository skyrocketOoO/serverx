package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type InviteUserIn struct {
	Email string `json:"email" validate:"required,email"`
}

func (u *Usecase) InviteUser(c context.Context, in InviteUserIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId: aws.String(u.cognitoSvc.UserPoolID),
		Username:   aws.String(in.Email),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(in.Email)},
			{Name: aws.String("email_verified"), Value: aws.String("true")},
		},
	}

	if _, err := u.cognitoSvc.Client.AdminCreateUser(c, input); err != nil {
		return er.W(err, er.BadRequest)
	}

	return nil
}
