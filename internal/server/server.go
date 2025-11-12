package server

import (
	"log"
	"net/http"

	"github.com/SidBroGG/dementia-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(h *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", h.RegisterHandler)
		r.Post("/login", h.LoginHandler)

		r.Post("/tasks", h.CreateTaskHandler)
		r.Get("/tasks", h.GetTaskHandler)
		r.Get("/tasks/{id}", h.GetTaskByIdHandler)
		r.Put("/tasks/{id}", h.UpdateTaskHandler)
		r.Delete("/tasks/{id}", h.DeleteTaskHandler)
	})

	log.Println("Router is configured")
	return r
}
