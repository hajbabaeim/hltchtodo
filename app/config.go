package app

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Name        string  `mapstructure:"name"`
	Version     float64 `mapstructure:"version"`
	Environment string  `mapstructure:"environment"`
	App         struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"app"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
	} `mapstructure:"database"`
	Logger struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logger"`
	SQS sqsConfig `mapstructure:"sqs"`
}

type sqsConfig struct {
	Region              string `mapstructure:"region"`
	AccessKeyID         string `mapstructure:"accessKeyId"`
	SecretAccessKey     string `mapstructure:"secretAccessKey"`
	QueueName           string `mapstructure:"queueName"`
	DeadLetterQueueName string `mapstructure:"deadLetterQueueName"`
	VisibilityTimeout   int    `mapstructure:"visibilityTimeout"`
	MaxRetries          int    `mapstructure:"maxRetries"`
	WaitTime            int    `mapstructure:"waitTime"`
	MaxMessages         int    `mapstructure:"maxMessages"`
	Endpoint            string `mapstructure:"endpoint"`
}

func getConfigFilename() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return fmt.Sprintf("cmd/configs/config-%s.yaml", env)
}

func (a *App) initConfig() error {
	filename := getConfigFilename()
	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	a.config = &cfg
	return nil
}
