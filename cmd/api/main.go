package main

import (
	"auth-template/internal/config"
	"auth-template/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	log.Printf("Server starting on port %s", cfg.ApiPort)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
