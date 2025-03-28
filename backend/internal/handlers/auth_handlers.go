package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
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
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid request body", nil)
		return
	}

	user, err := h.authService.Signin(input.Username, input.Password)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, false, "Invalid email or password", nil)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Failed to generate token", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Login successful", map[string]string{
		"user_id":      strconv.FormatInt(user.ID, 10),
		"access_token": token,
	})
}