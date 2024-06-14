package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiApiRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", BaseRoute)

	http.ListenAndServe(":8080", r)
}
