package usecase

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

func (uc *usecase) CreateItem(ctx context.Context, req *requests.CreateItemRequest) (*domain.TodoItem, error) {
	return nil, nil
}

func (uc *usecase) UpdateItem(ctx context.Context, req *requests.UpdateItemRequest) error {
	return nil
}

func (uc *usecase) DeleteItem(ctx context.Context, req *requests.DeleteItemRequest) error {
	return nil
}
