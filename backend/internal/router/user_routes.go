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
		midd.JWTMiddleware,
		midd.CORSMiddleware,
		midd.LoggingMiddleware,
		))
		mux.HandleFunc("GET /user/{id}", midd.ChainMiddleware(
		userHandler.GetUserHandlerById,
		midd.JWTMiddleware,
		midd.CORSMiddleware,
		midd.LoggingMiddleware,
		))
}
