package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

type updateItemDelivery struct {
	uc abstraction.Usecase
}

func UpdateTodoItem(c *gin.Context, uc abstraction.Usecase) (*domain.TodoItem, error) {
	h := &updateItemDelivery{uc}
	return h.handler(c)
}

func (d *updateItemDelivery) handler(c *gin.Context) (*domain.TodoItem, error) {
	ctx := c.Request.Context()
	var req requests.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	id := c.Param("id")
	req.Id = &id
	return d.uc.UpdateItem(ctx, &req)
}
