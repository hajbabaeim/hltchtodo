package usecase

import (
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/sirupsen/logrus"
)

type usecase struct {
	repo     abstraction.Repository
	sqs      SQSClientInterface
	queueURL string
	logger   *logrus.Logger
}

func NewUsecase(repo abstraction.Repository, sqs SQSClientInterface, queueURL string, logger *logrus.Logger) abstraction.Usecase {
	return &usecase{repo: repo, sqs: sqs, queueURL: queueURL, logger: logger}
}
