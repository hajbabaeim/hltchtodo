package delivery

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
)

type createItemDelivery struct {
	uc abstraction.Usecase
}

func CreateTodoItem(ctx context.Context, uc abstraction.Usecase) error {
	h := &createItemDelivery{uc: uc}
	return h.handler(ctx)
}

func (d *createItemDelivery) handler(ctx context.Context) error {
	return nil
}
