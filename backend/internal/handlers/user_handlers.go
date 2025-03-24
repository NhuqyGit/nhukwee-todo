package handlers

import (
	"backend/internal/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserHandlerById(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "user" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse id", http.StatusInternalServerError)
		return
	}
	user, err := h.service.GetUserByID(id)

	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
