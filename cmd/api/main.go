package main

import (
	"github.com/bsjhx/tournament-backend/internal/platform/db"
	"github.com/bsjhx/tournament-backend/internal/team/controllers"
	"github.com/bsjhx/tournament-backend/internal/team/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database := db.Init()
	db.RunMigrations()
	teamService := services.NewTeamService(database)

	router := gin.Default()

	controllers.RegisterTeamController(router, teamService)

	// healthcheck
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Fatal(router.Run(":8080"))
}
