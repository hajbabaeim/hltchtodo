package app

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Name        string
	Version     int
	Environment string
	App         struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}
	Logger struct {
		Level string `json:"level"`
	}
	Database postgresConfig
	SQS      sqsConfig
}

type sqsConfig struct {
	Region               string `yaml:"region" env:"AWS_REGION" env-default:"us-east-1"`
	AccessKeyID          string `yaml:"access_key_id" env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey      string `yaml:"secret_access_key" env:"AWS_SECRET_ACCESS_KEY"`
	QueueName            string `yaml:"queue_name" env:"SQS_QUEUE_NAME" env-default:"my-app-queue"`
	DeadLetterQueueName  string `yaml:"dlq_name" env:"SQS_DLQ_NAME" env-default:"my-app-dlq"`
	VisibilityTimeoutSec int32  `yaml:"visibility_timeout" env:"SQS_VISIBILITY_TIMEOUT" env-default:"30"`
	MaxRetries           int32  `yaml:"max_retries" env:"SQS_MAX_RETRIES" env-default:"3"`
	WaitTimeSeconds      int32  `yaml:"wait_time" env:"SQS_WAIT_TIME" env-default:"20"`
	MaxMessages          int32  `yaml:"max_messages" env:"SQS_MAX_MESSAGES" env-default:"10"`
	Endpoint             string `yaml:"endpoint" env:"SQS_ENDPOINT"` // For LocalStack
}

type postgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

func whichConfig() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}

func (a *App) initConfig() error {
	filename := fmt.Sprintf("cmd/configs/config-%s.json", whichConfig())
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var cfg config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return err
	}
	a.config = &cfg
	return nil
}
