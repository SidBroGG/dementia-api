package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SidBroGG/dementia-api/config"
	"github.com/SidBroGG/dementia-api/internal/server"
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

	// Router
	r := server.InitRouter()

	// Http server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting at localhost%s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
