package todo_item

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/repository"
	"github.com/hajbabaeim/hltchtodo/todo_item/usecase"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Module struct {
	UseCase abstraction.Usecase
	Repo    abstraction.Repository
}

func NewModule(db *gorm.DB, sqs *sqs.Client, queueURL string, logger *logrus.Logger) *Module {
	m := new(Module)
	m.Repo = repository.NewRepository(db)
	m.UseCase = usecase.NewUsecase(m.Repo, sqs, queueURL, logger)
	return m
}
