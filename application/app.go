package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: loadRoutes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("some thing went wrong : %s ", err)
	}
	return nil
}
