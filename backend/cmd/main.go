package main

import (
	"backend/internal/router"
	"log"
	"net/http"
)

func main() {
	// Init()
	server := router.NewRouter()

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
