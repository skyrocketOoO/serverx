package authucase

import (
	"github.com/skyrocketOoO/serverx/internal/service"
)

type Usecase struct {
	cognitoSvc *service.Cognito
}

func New(cognitoSvc *service.Cognito) *Usecase {
	return &Usecase{
		cognitoSvc: cognitoSvc,
	}
}
