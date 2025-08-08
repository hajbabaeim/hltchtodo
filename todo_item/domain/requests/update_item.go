package requests

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type UpdateItemRequest struct {
	Id          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}

func (r *UpdateItemRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}
