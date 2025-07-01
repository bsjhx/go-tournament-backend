package http

import (
	"github.com/bsjhx/tournament-backend/internal/team"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	r.GET("/teams/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		foundTeam, err := service.GetTeam(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if foundTeam.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
			return
		}

		c.JSON(http.StatusOK, foundTeam)
	})
}
