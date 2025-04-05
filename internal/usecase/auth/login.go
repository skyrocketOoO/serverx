package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type LoginIn struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type LoginOut struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	IDToken      string `json:"idToken"`
}

func (u *Usecase) Login(c context.Context, in LoginIn) (*LoginOut, error) {
	if err := validate.Get().Struct(in); err != nil {
		return nil, er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(u.cognitoSvc.ClientID),
		AuthParameters: map[string]string{
			"USERNAME":    in.Email,
			"PASSWORD":    in.Password,
			"SECRET_HASH": u.cognitoSvc.ComputeSecretHash(in.Email),
		},
	}

	resp, err := u.cognitoSvc.Client.InitiateAuth(c, input)
	if err != nil {
		return nil, er.W(err, er.Unauthorized)
	}

	if resp.ChallengeName == types.ChallengeNameTypeNewPasswordRequired {
		return nil, er.NewAppErr(er.NewPasswordRequired)
	}

	if resp.AuthenticationResult == nil {
		return nil, er.NewAppErr(er.Unknown)
	}

	return &LoginOut{
		AccessToken:  *resp.AuthenticationResult.AccessToken,
		RefreshToken: *resp.AuthenticationResult.RefreshToken,
		IDToken:      *resp.AuthenticationResult.IdToken,
	}, nil
}
