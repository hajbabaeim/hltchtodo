package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/helpers"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

type createItemDelivery struct {
	uc abstraction.Usecase
}

func CreateTodoItem(c *gin.Context, uc abstraction.Usecase) (*domain.TodoItem, error) {
	h := &createItemDelivery{uc: uc}
	return h.handler(c)
}

func (d *createItemDelivery) handler(c *gin.Context) (*domain.TodoItem, error) {
	ctx := c.Request.Context()
	body, err := c.Request.GetBody()
	if err != nil {
		return nil, err
	}
	req, err := helpers.Convert(body, new(requests.CreateItemRequest))
	if err != nil {
		return nil, err
	}
	return d.uc.CreateItem(ctx, req)
}
