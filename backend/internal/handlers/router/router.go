package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	RegisterUserRoutes(mux)
	return mux
}
