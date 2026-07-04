package main

import (
	"fmt"
	"log"
	"net/http"

	"uno/api"
)

func main() {

	mux := http.NewServeMux()

	// Register all API routes.
	api.RegisterRoutes(mux)

	addr := ":8080"

	fmt.Printf("UNO API Server running at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}