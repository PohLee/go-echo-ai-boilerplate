package main

import (
	"fmt"
	"log"

	"github.com/PohLee/go-echo-ai-boilerplate/pkg/config"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/database"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Connect to Database (using same logic as API)
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Running Migrations...")

	// 3. Register Models for Migration
	// TODO: Add models here as you create them
	// logic: models := []interface{}{ &model.User{}, &model.Product{} }

	err = db.AutoMigrate(
	// Models will go here
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations completed successfully!")
}
