package requests

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type CreateItemRequest struct {
	Description string    `json:"description" validate:"omitempty,alphaunicode"`
	DueDate     time.Time `json:"dueDate" validate:"omitempty,due_date_validator"`
}

func (r *CreateItemRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}
