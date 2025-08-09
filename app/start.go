package app

import (
	"context"
	"github.com/hajbabaeim/hltchtodo/server/http"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := http.NewServer(a.todoItemModule.UseCase)
	go func() {
		if err := server.Run(ctx); err != nil {
			a.logger.Errorf("failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	a.Stop()

}

func (a *App) Stop() {
	a.logger.Warnf("Stopping app")
}
