package app

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsc "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/sirupsen/logrus"
	"time"
)

type sqsClient struct {
	client   *sqs.Client
	queueURL string
	dlqURL   string
	config   sqsConfig
	logger   *logrus.Logger
}

func (a *App) initSQS() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Load AWS config
	cfg, err := a.loadAWSConfig(ctx)
	if err != nil {
		a.panicOnError(fmt.Errorf("failed to load AWS config: %w", err))
	}

	// Create SQS client
	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		if a.config.SQS.Endpoint != "" {
			o.BaseEndpoint = aws.String(a.config.SQS.Endpoint)
		}
	})

	// Initialize SQS wrapper
	a.sqs = &sqsClient{
		client: client,
		config: a.config.SQS,
		logger: a.logger,
	}

	// Setup queues
	if err := a.sqs.setupQueues(ctx); err != nil {
		a.panicOnError(fmt.Errorf("failed to setup SQS queues: %w", err))
	}

	// Health check
	if err := a.sqs.healthCheck(ctx); err != nil {
		a.panicOnError(fmt.Errorf("SQS health check failed: %w", err))
	}

	a.logger.WithFields(logrus.Fields{
		"queue_url": a.sqs.queueURL,
		"dlq_url":   a.sqs.dlqURL,
		"region":    a.config.SQS.Region,
	}).Info("SQS initialized successfully")
}

func (a *App) loadAWSConfig(ctx context.Context) (aws.Config, error) {
	// Load region + static creds (for LocalStack or testing)
	return awsc.LoadDefaultConfig(ctx,
		awsc.WithRegion(a.config.SQS.Region),
		awsc.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				a.config.SQS.AccessKeyID,
				a.config.SQS.SecretAccessKey,
				"",
			),
		),
	)
}

func (s *sqsClient) setupQueues(ctx context.Context) error {
	// Setup Dead Letter Queue first
	dlqURL, err := s.getOrCreateQueue(ctx, s.config.DeadLetterQueueName, nil)
	if err != nil {
		return fmt.Errorf("failed to setup DLQ: %w", err)
	}
	s.dlqURL = dlqURL

	// Get DLQ ARN for main queue redrive policy
	dlqArn, err := s.getQueueArn(ctx, dlqURL)
	if err != nil {
		return fmt.Errorf("failed to get DLQ ARN: %w", err)
	}

	// Setup main queue with DLQ redrive policy
	queueAttributes := map[string]string{
		string(types.QueueAttributeNameVisibilityTimeout):             fmt.Sprintf("%d", s.config.VisibilityTimeout),
		string(types.QueueAttributeNameReceiveMessageWaitTimeSeconds): fmt.Sprintf("%d", s.config.WaitTime),
		string(types.QueueAttributeNameRedrivePolicy):                 fmt.Sprintf(`{"deadLetterTargetArn":"%s","maxReceiveCount":%d}`, dlqArn, s.config.MaxRetries),
	}

	queueURL, err := s.getOrCreateQueue(ctx, s.config.QueueName, queueAttributes)
	if err != nil {
		return fmt.Errorf("failed to setup main queue: %w", err)
	}
	s.queueURL = queueURL

	return nil
}

func (s *sqsClient) getOrCreateQueue(ctx context.Context, queueName string, attributes map[string]string) (string, error) {
	// Try to get existing queue URL
	getQueueUrlInput := &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	}

	result, err := s.client.GetQueueUrl(ctx, getQueueUrlInput)
	if err == nil {
		s.logger.WithField("queue_name", queueName).Info("Found existing SQS queue")
		return *result.QueueUrl, nil
	}

	// Queue doesn't exist, create it
	s.logger.WithField("queue_name", queueName).Info("Creating new SQS queue")

	createQueueInput := &sqs.CreateQueueInput{
		QueueName: &queueName,
	}

	if attributes != nil {
		createQueueInput.Attributes = attributes
	}

	createResult, err := s.client.CreateQueue(ctx, createQueueInput)
	if err != nil {
		return "", fmt.Errorf("failed to create queue %s: %w", queueName, err)
	}

	return *createResult.QueueUrl, nil
}

func (s *sqsClient) getQueueArn(ctx context.Context, queueURL string) (string, error) {
	input := &sqs.GetQueueAttributesInput{
		QueueUrl: &queueURL,
		AttributeNames: []types.QueueAttributeName{
			types.QueueAttributeNameQueueArn,
		},
	}

	result, err := s.client.GetQueueAttributes(ctx, input)
	if err != nil {
		return "", err
	}

	arn, exists := result.Attributes[string(types.QueueAttributeNameQueueArn)]
	if !exists {
		return "", fmt.Errorf("queue ARN not found")
	}

	return arn, nil
}

func (s *sqsClient) healthCheck(ctx context.Context) error {
	input := &sqs.GetQueueAttributesInput{
		QueueUrl: &s.queueURL,
		AttributeNames: []types.QueueAttributeName{
			types.QueueAttributeNameApproximateNumberOfMessages,
		},
	}

	_, err := s.client.GetQueueAttributes(ctx, input)
	return err
}
