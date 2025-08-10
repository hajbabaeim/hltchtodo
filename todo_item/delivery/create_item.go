package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	var req requests.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	fmt.Printf(" + + + + + +the body: %+v\n", req)
	return d.uc.CreateItem(ctx, &req)
}
