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

func (r *repo) GetItemByUUID(uid string) (*domain.TodoItem, error) {
	var item domain.TodoItem
	uuidValue, err := uuid.Parse(uid)
	if err != nil {
		return nil, errors.New("uuid error")
	}
	err = r.db.Where("uuid=?", uuidValue).First(&item).Error
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

func (r *repo) DeleteItem(id uint64) error {
	return r.db.Delete(id).Error
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
