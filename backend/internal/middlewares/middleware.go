package middlewares

import (
	"backend/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var Username string = "nhuqy"
var Password string = "12345"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[LOG] %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("[LOG] Completed in %v", time.Since(start))
	}
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token doesn't exist", http.StatusUnauthorized)
			return
		}

		// Check format "Bearer <token>""
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Token with wrong format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Token invalid", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func WelcomeUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to my website")
}

func ChainMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// a(b) -> c(http.HandlerFunc) { b(w, r) }

//a(b) -> process(a) -> next(b)
//b(c) -> process(b) -> next(c)
