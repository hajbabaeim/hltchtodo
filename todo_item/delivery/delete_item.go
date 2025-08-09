package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/helpers"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

type deleteItemDelivery struct {
	uc abstraction.Usecase
}

func DeleteTodoItem(c *gin.Context, uc abstraction.Usecase) (bool, error) {
	h := &deleteItemDelivery{uc: uc}
	return h.handler(c)
}

func (d *deleteItemDelivery) handler(c *gin.Context) (bool, error) {
	ctx := c.Request.Context()
	req, err := helpers.Convert(c.Request.Body, new(requests.DeleteItemRequest))
	if err != nil {
		return false, err
	}
	return d.uc.DeleteItem(ctx, req)
}
