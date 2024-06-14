package application

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiApiRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", BaseRoute)
	fmt.Println("Running on Port 8080!")
	//! Open server Method 1

	//! Open server Method 2
	// err := http.ListenAndServe(":8080", router)

	return router
}
