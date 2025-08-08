package abstraction

import (
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
)

type Repository interface {
	CreateItem(item *domain.TodoItem) error
	GetItemByID(id uint64) (*domain.TodoItem, error)
	GetItemByUUID(uid string) (*domain.TodoItem, error)
	UpdateItem(item *domain.TodoItem) error
	DeleteItem(id uint64) error
	ListItems() ([]*domain.TodoItem, error)
}
