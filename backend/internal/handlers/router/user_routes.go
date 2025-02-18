package router

import (
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	userRepo := &repositories.UserRepository{}
	userServices := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userServices)

	mux.HandleFunc("GET /users", userHandler.GetUserHandler)
	mux.HandleFunc("GET /user/", userHandler.GetUserHandlerById)
}
