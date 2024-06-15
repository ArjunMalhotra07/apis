package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	redis  *redis.Client
}

func New() *App {
	app := &App{
		router: ChiApiRoutes(),
		redis:  redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) StartServer(c context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	err := a.redis.Ping(c).Err()
	if err != nil {
		return fmt.Errorf("can't connect to redis: %w", err)
	}
	defer func() {
		if err := a.redis.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()
	fmt.Println("Starting server")
	ch := make(chan error, 1)
	//! Method 1 to Start Server
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server %w", err)
		}
		close(ch)
	}()
	select {
	case err = <-ch:
		return err

	case <-c.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)

	}
	err = <-ch

	//! Method 2 to Start Server
	/*
		err := http.ListenAndServe(":8080", server.Handler)
		if err != nil {
			return fmt.Errorf("error %s", err)
		}
	*/
	return nil
}
