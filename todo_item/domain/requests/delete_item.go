package requests

import "github.com/go-playground/validator/v10"

type DeleteItemRequest struct {
	Id string `json:"id"`
}

func (r *DeleteItemRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}
