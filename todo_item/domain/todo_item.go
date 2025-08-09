package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TodoItem struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" gorm:"uuid"`
	Description string    `json:"description" gorm:"description"`
	DueDate     time.Time `json:"dueDate" gorm:"due_date"`
}
