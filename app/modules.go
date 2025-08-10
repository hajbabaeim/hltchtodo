package app

import (
	tdim "github.com/hajbabaeim/hltchtodo/todo_item"
)

func (a *App) initTodoItemModule() {
	a.todoItemModule = tdim.NewModule(a.postgres, a.sqs.client, a.sqs.queueURL, a.logger)
}
