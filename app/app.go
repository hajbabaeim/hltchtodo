package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	config     *config
	postgres   *gorm.DB
	logger     *logrus.Logger
	validator  *validator.Validate
	itemModule *td.Module
}

func NewApp() *App {
	a := new(App)
	if err := a.InitConfig(); err != nil {
		a.panicOnError(err)
	}
	return a
}

func (a *App) panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
