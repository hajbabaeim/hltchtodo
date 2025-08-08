package abstraction

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
)

type Repository interface {
	CreateItem(ctx context.Context, item *domain.TodoItem) (*domain.TodoItem, error)
	UpdateItem(ctx context.Context, item *domain.TodoItem) error
	DeleteItem(ctx context.Context, id int) error
}
