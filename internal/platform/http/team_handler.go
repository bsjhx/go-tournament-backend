package http

import (
	"github.com/bsjhx/tournament-backend/internal/team"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterTeamRoutes(r *gin.Engine, service *team.Service) {
	r.POST("/teams", func(context *gin.Context) {
		var input struct {
			Name string `json:"name"`
		}
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		t := service.AddTeam(input.Name)
		context.JSON(http.StatusOK, t)
	})

	r.GET("/teams", func(context *gin.Context) {
		context.JSON(http.StatusOK, service.ListTeams())
	})
}
