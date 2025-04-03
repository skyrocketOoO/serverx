package cognito

// import (
// 	"context"
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"errors"

// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
// 	"github.com/skyrocketOoO/erx/erx"
// )

// var (
// 	cognitoClient *cognitoidentityprovider.Client
// 	UserPoolID    string
// 	ClientID      string
// 	ClientSecret  string
// )

// func Get() *cognitoidentityprovider.Client {
// 	return cognitoClient
// }

// func New() error {
// 	UserPoolID = local.GetEnv("COGNITO_USERPOOLID")
// 	ClientID = local.GetEnv("COGNITO_CLIENTID")
// 	ClientSecret = local.GetEnv("COGNITO_CLIENTSECRET")

// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion(local.GetEnv("COGNITO_REGION")), // 仍然可以從環境變數獲取區域
// 	)
// 	if err != nil {
// 		return erx.W(err, "Error loading AWS config")
// 	}

// 	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
// 	return nil
// }

// func GenerateSecretHash(input string) string {
// 	h := hmac.New(sha256.New, []byte(ClientSecret))
// 	h.Write([]byte(input + ClientID)) // clientID is your Cognito App Client ID
// 	return base64.StdEncoding.EncodeToString(h.Sum(nil))
// }

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
