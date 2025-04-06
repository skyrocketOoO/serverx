package authucase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	validate "github.com/skyrocketOoO/serverx/internal/service/validator"
)

type SetNewPasswordIn struct {
	Email   string `json:"email"   validate:"required"`
	OldPass string `json:"oldPass" validate:"required"`
	NewPass string `json:"newPass" validate:"required"`
}

func (u *Usecase) SetNewPassword(c context.Context, in SetNewPasswordIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	// Step 1: InitiateAuth with temp password
	authResp, err := u.cognitoSvc.Client.InitiateAuth(c,
		&cognitoidentityprovider.InitiateAuthInput{
			AuthFlow: "USER_PASSWORD_AUTH",
			ClientId: aws.String(u.cognitoSvc.ClientID),
			AuthParameters: map[string]string{
				"USERNAME": in.Email,
				"PASSWORD": in.OldPass,
			},
		},
	)
	if err != nil {
		return er.W(err)
	}

	// Step 2: Check if response requires NEW_PASSWORD_REQUIRED challenge
	if authResp.ChallengeName != types.ChallengeNameTypeNewPasswordRequired {
		return er.NewAppErr(er.AlreadyResetOTP)
	}

	// Step 3: Respond to the challenge with new password
	if _, err = u.cognitoSvc.Client.RespondToAuthChallenge(
		c,
		&cognitoidentityprovider.RespondToAuthChallengeInput{
			ClientId:      aws.String(u.cognitoSvc.ClientID),
			ChallengeName: types.ChallengeNameTypeNewPasswordRequired,
			ChallengeResponses: map[string]string{
				"USERNAME":     in.Email,
				"NEW_PASSWORD": in.NewPass,
			},
			Session: authResp.Session,
		},
	); err != nil {
		return er.W(err)
	}

	return nil
}
