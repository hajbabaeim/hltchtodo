package usecase

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSClientInterface interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type SQSClientWrapper struct {
	client *sqs.Client
}

func NewSQSClientWrapper(client *sqs.Client) SQSClientInterface {
	return &SQSClientWrapper{client: client}
}

func (w *SQSClientWrapper) SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {
	return w.client.SendMessage(ctx, params, optFns...)
}
