package abstraction

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

type Usecase interface {
	CreateItem(ctx context.Context, req *requests.CreateItemRequest) (*domain.TodoItem, error)
	GetItem(ctx context.Context, req *requests.GetItemRequest) (*domain.TodoItem, error)
	UpdateItem(ctx context.Context, req *requests.UpdateItemRequest) (*domain.TodoItem, error)
	DeleteItem(ctx context.Context, req *requests.DeleteItemRequest) (bool, error)
	ListItems(ctx context.Context) ([]*domain.TodoItem, error)
}
