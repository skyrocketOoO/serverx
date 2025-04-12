package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSService struct {
	client *sns.Client
	topic  string
}

func NewSNS(ctx context.Context, topicArn string) (*SNSService, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &SNSService{
		client: sns.NewFromConfig(cfg),
		topic:  topicArn,
	}, nil
}

func (s *SNSService) Publish(ctx context.Context, msg string) error {
	_, err := s.client.Publish(ctx, &sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: aws.String(s.topic),
	})
	return err
}
