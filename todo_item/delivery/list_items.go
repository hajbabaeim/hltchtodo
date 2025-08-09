package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
)

type listItemsDelivery struct {
	uc abstraction.Usecase
}

func ListItems(c *gin.Context, uc abstraction.Usecase) ([]*domain.TodoItem, error) {
	h := &listItemsDelivery{uc}
	return h.handler(c)
}

func (d *listItemsDelivery) handler(c *gin.Context) ([]*domain.TodoItem, error) {
	ctx := c.Request.Context()
	return d.uc.ListItems(ctx)
}
