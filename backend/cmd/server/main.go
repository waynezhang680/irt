package main

import (
	"log"

	"backend/internal/infrastructure/persistence"
	"backend/internal/interfaces/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := persistence.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	routes.SetupAuthRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
