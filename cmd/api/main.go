package main

import (
	"github.com/bsjhx/tournament-backend/internal/platform/db"
	"github.com/bsjhx/tournament-backend/internal/platform/http"
	"github.com/bsjhx/tournament-backend/internal/team"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.RunMigrations()
	database := db.Init()
	teamService := team.NewService(database)

	router := gin.Default()

	http.RegisterTeamRoutes(router, teamService)

	// healthcheck
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Fatal(router.Run(":8080"))
}
