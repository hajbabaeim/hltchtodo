package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hajbabaeim/hltchtodo/server/http"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Start() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	server := http.NewServer(a.todoItemModule.UseCase)
	server.SetupRoutes()
	router := server.GetRouter()
	go func(r *gin.Engine) {
		if err := r.Run(fmt.Sprintf(":%d", a.config.App.Port)); err != nil {
			a.logger.Errorf("failed to start server: %v", err)
		}
	}(router)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	a.Stop()

}

func (a *App) Stop() {
	a.logger.Warnf("Stopping app")
}
