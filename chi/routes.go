package chi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiApiRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", BaseRoute)
	fmt.Println("Running on Port 8080!")
	//! Open server Method 1
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	//! Open server Method 2
	// err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
}
