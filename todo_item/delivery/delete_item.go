package delivery

import (
	"github.com/gin-gonic/gin"
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
	req := new(requests.DeleteItemRequest)
	req.Id = c.Param("id")
	return d.uc.DeleteItem(ctx, req)
}
