package chi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiApiRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", BaseRoute)
	fmt.Println("Running on Port 8080!")
	http.ListenAndServe(":8080", r)
}
