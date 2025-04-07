package authucase

import (
	"github.com/skyrocketOoO/serverx/internal/service/aws"
)

type Usecase struct {
	cognitoSvc *aws.Cognito
}

func New(cognitoSvc *aws.Cognito) *Usecase {
	return &Usecase{
		cognitoSvc: cognitoSvc,
	}
}
