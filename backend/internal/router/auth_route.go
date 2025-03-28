package router

import (
	"backend/internal/handlers"
	midd "backend/internal/middlewares"
	"backend/internal/repositories"
	"backend/internal/services"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	mux.HandleFunc("/signup", midd.ChainMiddleware(
		authHandler.SignupHandler,
		midd.CORSMiddleware,
	))
	mux.HandleFunc("/signin", midd.ChainMiddleware(
		authHandler.SigninHandler,
		midd.CORSMiddleware,
	))
}
