package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSService struct {
	client *sqs.Client
	queue  string
}

func NewSQS(ctx context.Context, queueURL string) (*SQSService, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &SQSService{
		client: sqs.NewFromConfig(cfg),
		queue:  queueURL,
	}, nil
}

func (s *SQSService) SendMessage(ctx context.Context, msg string) error {
	_, err := s.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.queue),
		MessageBody: aws.String(msg),
	})
	return err
}

func (s *SQSService) ReceiveAndDelete(ctx context.Context, handler func(string)) error {
	output, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(s.queue),
		MaxNumberOfMessages: 5,
		WaitTimeSeconds:     10,
	})
	if err != nil {
		return err
	}

	for _, msg := range output.Messages {
		handler(*msg.Body)

		// delete after handling
		_, err := s.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(s.queue),
			ReceiptHandle: msg.ReceiptHandle,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
