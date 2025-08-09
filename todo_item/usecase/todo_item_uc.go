package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/hajbabaeim/hltchtodo/helpers"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

func (uc *usecase) CreateItem(ctx context.Context, req *requests.CreateItemRequest) (*domain.TodoItem, error) {
	item, err := helpers.Convert(req, new(domain.TodoItem))
	if err != nil {
		return nil, err
	}
	uid := uuid.New()
	item.UUID = uid
	if err = uc.repo.CreateItem(item); err != nil {
		return nil, err
	}
	return uc.repo.GetItemByUUID(item.UUID)
}

func (uc *usecase) UpdateItem(ctx context.Context, req *requests.UpdateItemRequest) (*domain.TodoItem, error) {
	if req.Id == nil && *req.Id == 0 {
		return nil, errors.New("id is required")
	}
	oldItem, err := uc.repo.GetItemByID(*req.Id)
	if err != nil {
		return nil, err
	}
	if req.Description != nil && oldItem.Description != *req.Description {
		oldItem.Description = *req.Description
	}
	if req.DueDate != nil && !oldItem.DueDate.Equal(*req.DueDate) {
		oldItem.DueDate = *req.DueDate
	}
	if err = uc.repo.UpdateItem(oldItem); err != nil {
		return nil, err
	}
	return uc.repo.GetItemByUUID(oldItem.UUID)
}

func (uc *usecase) DeleteItem(ctx context.Context, req *requests.DeleteItemRequest) (bool, error) {
	if req.Id == "" {
		return false, errors.New("id is required")
	}
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		return false, err
	}
	if err = uc.repo.DeleteItem(uid); err != nil {
		return false, err
	}
	return true, nil
}

func (uc *usecase) ListItems(ctx context.Context) ([]*domain.TodoItem, error) {
	items, err := uc.repo.ListItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *usecase) GetItem(ctx context.Context, req *requests.GetItemRequest) (*domain.TodoItem, error) {
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetItemByUUID(uid)
}
