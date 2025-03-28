package router

import (
	"backend/internal/handlers"
	midd "backend/internal/middlewares"
	"backend/internal/repositories"
	"backend/internal/services"
	"fmt"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	userRepo := repositories.NewUserRepository()
	fmt.Println(userRepo)
	userServices := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userServices)

	mux.HandleFunc("GET /users", midd.ChainMiddleware(
		userHandler.GetUserHandler,
		midd.LoggingMiddleware,
		midd.JWTMiddleware,
		midd.CORSMiddleware,
		))
		mux.HandleFunc("GET /user/{id}", midd.ChainMiddleware(
		userHandler.GetUserHandlerById,
		midd.LoggingMiddleware,
		midd.JWTMiddleware,
		midd.CORSMiddleware,
		))
}
