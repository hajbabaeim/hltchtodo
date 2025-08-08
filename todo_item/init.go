package todo_item

import (
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/repository"
	"github.com/hajbabaeim/hltchtodo/todo_item/usecase"
	"gorm.io/gorm"
)

type Module struct {
	UseCase abstraction.Usecase
	Repo    abstraction.Repository
}

func NewModule(db *gorm.DB) *Module {
	m := new(Module)
	m.Repo = repository.NewRepository(db)
	m.UseCase = usecase.NewUsecase(m.Repo)
	return m
}
