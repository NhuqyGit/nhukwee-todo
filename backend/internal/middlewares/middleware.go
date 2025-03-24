package middlewares

import (
	"fmt"
	"log"
	"net/http"
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

func AuthLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// username, password, ok := r.BasicAuth()
		if Username != "nhuqy" || Password != "12345" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
