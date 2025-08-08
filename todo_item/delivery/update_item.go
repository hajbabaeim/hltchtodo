package delivery

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
)

type updateItemDelivery struct {
	uc abstraction.Usecase
}

func UpdateTodoItem(ctx context.Context, uc abstraction.Usecase) error {
	h := &updateItemDelivery{uc}
	return h.handler(ctx)
}

func (d *updateItemDelivery) handler(ctx context.Context) error {
	return nil
}
