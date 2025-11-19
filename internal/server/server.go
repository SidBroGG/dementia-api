package server

import (
	"log"
	"net/http"

	"github.com/SidBroGG/dementia-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(authH *handlers.AuthHandler, taskH *handlers.TaskHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", authH.RegisterHandler)
		r.Post("/login", authH.LoginHandler)

		r.Post("/tasks", taskH.CreateTaskHandler)
		r.Get("/tasks", taskH.GetTaskHandler)
		r.Get("/tasks/{id}", taskH.GetTaskByIdHandler)
		r.Put("/tasks/{id}", taskH.UpdateTaskHandler)
		r.Delete("/tasks/{id}", taskH.DeleteTaskHandler)
	})

	log.Println("Router is configured")
	return r
}
