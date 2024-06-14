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
	//! Method 1 to Start Server
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	//! Method 2 to Start Server
	/*
		err := http.ListenAndServe(":8080", server.Handler)
		if err != nil {
			return fmt.Errorf("error %s", err)
		}
	*/
	return nil
}
