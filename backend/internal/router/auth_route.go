package router

import (
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	mux.HandleFunc("/signup", authHandler.SignupHandler)
	mux.HandleFunc("/signin", authHandler.SigninHandler)
}
