package authucase

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
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
		return nil, err
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
		return nil, fmt.Errorf("failed to initiate auth: %w", err)
	}

	if resp.ChallengeName == types.ChallengeNameTypeNewPasswordRequired {
	}

	if resp.AuthenticationResult == nil {
		return nil, fmt.Errorf("authentication result is nil")
	}

	return &LoginOut{
		AccessToken:  *resp.AuthenticationResult.AccessToken,
		RefreshToken: *resp.AuthenticationResult.RefreshToken,
		IDToken:      *resp.AuthenticationResult.IdToken,
	}, nil
}
