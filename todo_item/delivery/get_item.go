package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
)

type getItemDelivery struct {
	uc abstraction.Usecase
}

func GetTodoItem(c *gin.Context, uc abstraction.Usecase) (*domain.TodoItem, error) {
	h := &getItemDelivery{uc: uc}
	return h.handler(c)
}

func (d *getItemDelivery) handler(c *gin.Context) (*domain.TodoItem, error) {
	ctx := c.Request.Context()
	req := new(requests.GetItemRequest)
	req.Id = c.Param("id")
	return d.uc.GetItem(ctx, req)
}
