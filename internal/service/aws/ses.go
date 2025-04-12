package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type SESService struct {
	client *ses.Client
	from   string
}

func NewSES(ctx context.Context, senderEmail string) (*SESService, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &SESService{
		client: ses.NewFromConfig(cfg),
		from:   senderEmail,
	}, nil
}

func (s *SESService) SendEmail(
	ctx context.Context,
	toEmail, subject, bodyHTML, bodyText string,
) error {
	input := &ses.SendEmailInput{
		Source: aws.String(s.from),
		Destination: &types.Destination{
			ToAddresses: []string{toEmail},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: aws.String(subject),
			},
			Body: &types.Body{
				Html: &types.Content{
					Data: aws.String(bodyHTML),
				},
				Text: &types.Content{
					Data: aws.String(bodyText),
				},
			},
		},
	}

	_, err := s.client.SendEmail(ctx, input)
	return err
}
