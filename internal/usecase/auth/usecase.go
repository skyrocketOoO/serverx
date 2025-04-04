package authusecase

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type Usecase struct {
	cognitoCli *cognitoidentityprovider.Client
}

func New(cognitoCli *cognitoidentityprovider.Client) *Usecase {
	return &Usecase{
		cognitoCli: cognitoCli,
	}
}
