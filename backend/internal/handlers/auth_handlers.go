package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	log.Println(r.Method)
	if r.Method != http.MethodPost {
		log.Println("err", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("HELLL")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.authService.Signup(&user); err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (h *AuthHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Method:", r.Method)

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Đọc raw body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Debug: In raw body để kiểm tra JSON
	fmt.Println("Raw Body:", string(body))

	// Giải mã JSON
	if err := json.Unmarshal(body, &input); err != nil {
		fmt.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Kiểm tra dữ liệu
	fmt.Println("Parsed Username:", input.Username)
	fmt.Println("Parsed Password:", input.Password)

	// if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	user, err := h.authService.Signin(input.Username, input.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "user_id": strconv.FormatInt(user.ID, 10)})
}
