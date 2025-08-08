package abstraction

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

type Usecase interface {
	CreateItem(ctx context.Context, req *requests.CreateItemRequest) (*domain.TodoItem, error)
	UpdateItem(ctx context.Context, req *requests.UpdateItemRequest) error
	DeleteItem(ctx context.Context, req *requests.DeleteItemRequest) error
}
