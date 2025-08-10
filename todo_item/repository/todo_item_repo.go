package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"gorm.io/gorm"
)

func (r *repo) CreateItem(item *domain.TodoItem) error {
	return r.db.Create(item).Error
}

func (r *repo) GetItemByID(id uint64) (*domain.TodoItem, error) {
	var item domain.TodoItem
	err := r.db.First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, errors.New("db error")
	}
	return &item, nil
}

func (r *repo) GetItemByUUID(uid uuid.UUID) (*domain.TodoItem, error) {
	var item domain.TodoItem
	err := r.db.Where("uuid=?", uid).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, errors.New("db error")
	}
	return &item, nil
}

func (r *repo) UpdateItem(item *domain.TodoItem) error {
	return r.db.Save(item).Error
}

func (r *repo) DeleteItem(id uuid.UUID) error {
	return r.db.Where("uuid=?", id).Delete(&domain.TodoItem{}).Error
}

func (r *repo) ListItems() ([]*domain.TodoItem, error) {
	var items []*domain.TodoItem
	err := r.db.Find(&items).Error
	if err != nil {
		return nil, errors.New("db error")
	}
	return items, nil
}

func (r *repo) ListItemsByUserId(userId uint64) ([]*domain.TodoItem, error) {
	var items []*domain.TodoItem
	err := r.db.Where("user_id=?", userId).Find(&items).Error
	if err != nil {
		return nil, errors.New("db error")
	}
	return items, nil
}
