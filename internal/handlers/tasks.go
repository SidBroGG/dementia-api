package handlers

import (
	"fmt"
	"net/http"

	"github.com/SidBroGG/dementia-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	svc *service.AuthService
}

func NewTaskHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Task created")
}

func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Get task handler ok")
}

func (h *TaskHandler) GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get task by id handler ok (id: %v)", taskId)
}

func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task updated ok (id: %v)", taskId)
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task deleted (id: %v)", taskId)
}
