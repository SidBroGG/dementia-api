package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Task created")
}

func (h *Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Get task handler ok")
}

func (h *Handler) GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get task by id handler ok (id: %v)", taskId)
}

func (h *Handler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task updated ok (id: %v)", taskId)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task deleted (id: %v)", taskId)
}
