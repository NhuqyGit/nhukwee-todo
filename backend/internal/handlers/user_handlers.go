package handlers

import (
	"backend/internal/services"
	"encoding/json"
	"fmt"
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
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO")
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "users" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	user, err := h.service.GetUserByID(id)

	fmt.Println("HELLO")
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
