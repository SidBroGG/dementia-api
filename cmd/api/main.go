package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SidBroGG/dementia-api/config"
	"github.com/SidBroGG/dementia-api/internal/auth"
	"github.com/SidBroGG/dementia-api/internal/handlers"
	"github.com/SidBroGG/dementia-api/internal/server"
	"github.com/SidBroGG/dementia-api/internal/service"
	"github.com/SidBroGG/dementia-api/internal/store"
)

func main() {
	// Config
	cfg := config.LoadConfig()
	log.Printf("Loaded config:\nPort: %v\n", cfg.Port)

	// DB
	db, err := store.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer db.Close()
	storeRepo := store.NewStore(db)

	// Auth (JWT)
	jwtKey := []byte(cfg.JWTSecret)
	auth := auth.NewJWTAuth(jwtKey, cfg.TokenTTL)

	// Service
	svc := service.NewService(*storeRepo, auth)

	// Handlers handler
	h := handlers.New(svc)

	// Router
	r := server.InitRouter(h)

	// Http server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting at localhost%s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
