package server

import (
	"log"
	"net/http"

	"github.com/SidBroGG/dementia-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", handlers.RegisterHandler)
		r.Post("/login", handlers.LoginHandler)

		r.Post("/tasks", handlers.CreateTaskHandler)
		r.Get("/tasks", handlers.GetTaskHandler)
		r.Get("/tasks/{id}", handlers.GetTaskByIdHandler)
		r.Put("/tasks/{id}", handlers.UpdateTaskHandler)
		r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)
	})

	log.Println("Router is configured")
	return r
}
