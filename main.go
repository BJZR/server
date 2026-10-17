package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bjzr/server/config"
	"github.com/bjzr/server/handlers"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/health", handlers.Health)

	addr := "0.0.0.0:" + cfg.Port
	fmt.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
