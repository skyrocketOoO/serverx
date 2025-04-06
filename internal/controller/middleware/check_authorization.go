package middleware

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	"github.com/skyrocketOoO/serverx/internal/service"
)

func CheckAuthorization(cognito *service.Cognito) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			er.Bind(c, er.NewAppErr(er.MissingAuthorizationHeader))
			return
		}

		const bearerPrefix = "Bearer "
		token = strings.TrimPrefix(token, bearerPrefix)

		email, err := getUserMail(c, cognito, token)
		if err != nil {
			er.Bind(c, er.W(err, er.Unauthorized))
			return
		}

		c.Set("email", email)

		c.Next()
	}
}

func getUserMail(c context.Context, cognito *service.Cognito, token string) (string, error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: &token,
	}

	resp, err := cognito.Client.GetUser(c, input)
	if err != nil {
		return "", er.W(err, er.Unauthorized)
	}

	// Extract the email from the response
	for _, attr := range resp.UserAttributes {
		if attr.Name != nil && *attr.Name == "email" && attr.Value != nil {
			return *attr.Value, nil
		}
	}

	return "", er.NewAppErr(er.Unauthorized)
}
