package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SidBroGG/dementia-api/internal/model"
	"github.com/SidBroGG/dementia-api/internal/service"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Handler struct {
	svc *service.AuthService
}

func New(svc *service.AuthService) *Handler {
	return &Handler{svc: svc}
}

func parseAuthRequest(r *http.Request) (*model.AuthRequest, error) {
	defer r.Body.Close()

	data := &model.AuthRequest{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := parseAuthRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(data); err != nil {
		http.Error(w, "Register data format error", http.StatusBadRequest)
		return
	}

	if err := h.svc.Register(ctx, *data); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Successfully registered")
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := parseAuthRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(data); err != nil {
		http.Error(w, "Login data format error", http.StatusBadRequest)
		return
	}

	resp, err := h.svc.Login(ctx, *data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
