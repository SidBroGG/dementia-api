package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type LoginData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
}

func parseLoginData(r *http.Request) (*LoginData, error) {
	defer r.Body.Close()

	data := &LoginData{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseLoginData(r)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = validate.Struct(data)
	if err != nil {
		http.Error(w, "Register data format error", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Successfully registered")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseLoginData(r)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = validate.Struct(data)
	if err != nil {
		http.Error(w, "Login data format error", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Successfully logined in")
}
