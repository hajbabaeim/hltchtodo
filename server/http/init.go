package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/todo_item/abstraction"
	"github.com/hajbabaeim/hltchtodo/todo_item/delivery"
	"net/http"
)

type server struct {
	todoUc abstraction.Usecase
	router *gin.Engine
}

func NewServer(todoUsecase abstraction.Usecase) *server {
	return &server{
		todoUc: todoUsecase,
		router: gin.Default(),
	}
}

func (s *server) Run(ctx context.Context) error {
	return nil
}

func (s *server) GetRouter() *gin.Engine {
	return s.router
}

func (s *server) SetupRoutes() {
	// handling check health requests
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "todo-item",
		})
	})
	// handling other requests
	api := s.router.Group("/api/v1")
	{
		todo := api.Group("/todos")
		{
			todo.POST("", func(c *gin.Context) {
				result, err := delivery.CreateTodoItem(c, s.todoUc)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusOK, result)
			})
			todo.GET("/:id", func(c *gin.Context) {
				result, err := delivery.GetTodoItem(c, s.todoUc)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusOK, result)
			})
			todo.PUT("/:id", func(c *gin.Context) {
				result, err := delivery.UpdateTodoItem(c, s.todoUc)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusOK, result)
			})
			todo.DELETE("/:id", func(c *gin.Context) {
				result, err := delivery.DeleteTodoItem(c, s.todoUc)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusOK, result)
			})
			todo.GET("", func(c *gin.Context) {
				result, err := delivery.ListItems(c, s.todoUc)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusOK, result)
			})
		}
	}
}
