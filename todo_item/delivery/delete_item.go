package delivery

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
)

type deleteItemDelivery struct {
	uc abstraction.Usecase
}

func DeleteTodoItem(ctx context.Context, uc abstraction.Usecase) error {
	h := &deleteItemDelivery{uc: uc}
	return h.handler(ctx)
}

func (d *deleteItemDelivery) handler(ctx context.Context) error {
	return nil
}
