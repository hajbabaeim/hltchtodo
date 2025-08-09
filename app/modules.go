package app

import (
	tdim "github.com/hajbabaeim/hltchtodo/todo_item"
)

func (a *App) initTodoItemModule() {
	a.todoItemModule = tdim.NewModule(a.postgres)
}
