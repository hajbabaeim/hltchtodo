package requests

import (
	"github.com/go-playground/validator/v10"
)

type UpdateItemRequest struct {
	Id          *string `json:"id"`
	Description *string `json:"description"`
	DueDate     *string `json:"due_date"`
}

func (r *UpdateItemRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}
