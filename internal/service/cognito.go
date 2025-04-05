package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
)

type Cognito struct {
	Client       *cognitoidentityprovider.Client
	UserPoolID   string
	ClientID     string
	ClientSecret string
}

func NewCognito(c context.Context) (*Cognito, error) {
	cfg, err := config.LoadDefaultConfig(c)
	if err != nil {
		return nil, er.W(err, "Error loading AWS config")
	}

	return &Cognito{
		Client:       cognitoidentityprovider.NewFromConfig(cfg),
		UserPoolID:   os.Getenv("COGNITO_USERPOOLID"),
		ClientID:     os.Getenv("COGNITO_CLIENTID"),
		ClientSecret: os.Getenv("COGNITO_CLIENTSECRET"),
	}, nil
}

func (c *Cognito) ComputeSecretHash(email string) string {
	h := hmac.New(sha256.New, []byte(c.ClientSecret))
	h.Write([]byte(email + c.ClientID))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// func ValidateTokenWithCognito(c context.Context, token string) (string, error) {
// 	// Call Cognito to validate the token
// 	input := &cognitoidentityprovider.GetUserInput{
// 		AccessToken: &token,
// 	}

// 	resp, err := cognitoClient.GetUser(c, input)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Extract the email from the response
// 	for _, attr := range resp.UserAttributes {
// 		if *attr.Name == "email" {
// 			return *attr.Value, nil
// 		}
// 	}

// 	return "", errors.New("email not found in token")
// }
