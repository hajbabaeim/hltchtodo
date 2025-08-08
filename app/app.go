package app

import (
	"github.com/go-playground/validator/v10"
	tdm "github.com/hajbabaeim/hltchtodo/todo_item"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	config         *config
	postgres       *gorm.DB
	logger         *logrus.Logger
	validator      *validator.Validate
	todoItemModule *tdm.Module
	sqs            *SQSClient
	//httpServer     *http.Server
}

func NewApp() *App {
	a := new(App)
	if err := a.initConfig(); err != nil {
		a.panicOnError(err)
	}
	return a
}

func (a *App) panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *App) Init() {
	a.initLogger()
	a.initValidator()
	a.initPostgres()
	a.initSQS()
}

func (a *App) InitModules() {
	a.initTodoItemModule()
}
