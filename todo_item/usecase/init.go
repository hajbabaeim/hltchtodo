package usecase

import "github.com/hajbabaeim/hltchtodo/todo_item/abstraction"

type usecase struct {
	repo abstraction.Repository
}

func NewUsecase(repo abstraction.Repository) abstraction.Usecase {
	return &usecase{repo: repo}
}
