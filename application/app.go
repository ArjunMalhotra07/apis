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
		router: ChiApiRoutes(),
	}
	return app
}

func (a *App) StartServer(c context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	return nil
}
