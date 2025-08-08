package requests

import "time"

type UpdateItemRequest struct {
	Id          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}
