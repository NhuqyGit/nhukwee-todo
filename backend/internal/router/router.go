package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// middleHandler := midd.ChainMiddleware(midd.WelcomeUser,
	// 	midd.LoggingMiddleware,
	// 	midd.AuthLogin,
	// )

	// mux.HandleFunc("/", middleHandler)

	RegisterAuthRoutes(mux)
	RegisterUserRoutes(mux)
	return mux
}
