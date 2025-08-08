package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TodoItem struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" db:"uuid"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DueDate     time.Time `json:"dueDate" db:"due_date"`
}
