package aws

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type CWLogWriter struct {
	client        *cloudwatchlogs.Client
	logGroupName  string
	logStreamName string
	sequenceToken *string
	mu            sync.Mutex
}

func NewCWLogWriter(ctx context.Context, groupName, streamName string) (*CWLogWriter, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := cloudwatchlogs.NewFromConfig(cfg)

	// Ensure log group and stream exist
	_, _ = client.CreateLogGroup(ctx, &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(groupName),
	})

	_, _ = client.CreateLogStream(ctx, &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(groupName),
		LogStreamName: aws.String(streamName),
	})

	return &CWLogWriter{
		client:        client,
		logGroupName:  groupName,
		logStreamName: streamName,
	}, nil
}

func (w *CWLogWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	input := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []types.InputLogEvent{
			{
				Message:   aws.String(string(bytes.TrimSpace(p))),
				Timestamp: aws.Int64(time.Now().UnixMilli()),
			},
		},
		LogGroupName:  aws.String(w.logGroupName),
		LogStreamName: aws.String(w.logStreamName),
		SequenceToken: w.sequenceToken,
	}

	out, err := w.client.PutLogEvents(context.TODO(), input)
	if err != nil {
		return 0, err
	}

	w.sequenceToken = out.NextSequenceToken
	return len(p), nil
}
