package abstraction

import (
	"github.com/google/uuid"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
)

type Repository interface {
	CreateItem(item *domain.TodoItem) error
	GetItemByID(id uint64) (*domain.TodoItem, error)
	GetItemByUUID(uid uuid.UUID) (*domain.TodoItem, error)
	UpdateItem(item *domain.TodoItem) error
	DeleteItem(id uuid.UUID) error
	ListItems() ([]*domain.TodoItem, error)
}
