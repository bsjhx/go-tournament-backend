package http

import (
	"github.com/bsjhx/tournament-backend/internal/team"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterTeamRoutes(r *gin.Engine, service *team.Service) {
	r.POST("/teams", func(c *gin.Context) {
		var dto team.CreateTeamDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdTeam, err := service.AddTeam(dto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, createdTeam)
	})

	r.GET("/teams", func(c *gin.Context) {
		teams, _ := service.ListTeams()
		c.JSON(http.StatusOK, teams)
	})
}
