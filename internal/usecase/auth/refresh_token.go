package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type RefreshTokenIn struct {
	Email        string `json:"email"        validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokenOut struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	IDToken      string `json:"idToken"`
}

func (u *Usecase) RefreshToken(c context.Context, in RefreshTokenIn) (*RefreshTokenOut, error) {
	if err := validate.Get().Struct(in); err != nil {
		return nil, er.W(err, er.ValidateInput)
	}

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeRefreshTokenAuth,
		ClientId: aws.String(u.cognitoSvc.ClientID),
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": in.RefreshToken,
			"SECRET_HASH":   u.cognitoSvc.ComputeSecretHash(in.Email),
		},
	}

	resp, err := u.cognitoSvc.Client.InitiateAuth(c, input)
	if err != nil {
		return nil, er.W(err)
	}

	return &RefreshTokenOut{
		AccessToken:  *resp.AuthenticationResult.AccessToken,
		RefreshToken: *resp.AuthenticationResult.RefreshToken,
		IDToken:      *resp.AuthenticationResult.IdToken,
	}, nil
}
