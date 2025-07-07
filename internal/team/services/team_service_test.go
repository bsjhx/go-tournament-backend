package services

import (
	"github.com/bsjhx/tournament-backend/internal/platform/db"
	"os"
	"testing"

	"github.com/bsjhx/tournament-backend/internal/team/model"
)

const testDBPath = "tournament.db"

func teardownTestDB() {
	_ = os.Remove(testDBPath)
}

func TestTeamService_AddAndGetTeam(t *testing.T) {
	testDB := db.Init()
	db.RunMigrations()
	defer teardownTestDB()

	service := NewTeamService(testDB)

	country := model.Country{
		Name:       "Testland",
		Shortcut:   "TST",
		Federation: "E",
	}
	if err := testDB.Create(&country).Error; err != nil {
		t.Fatalf("failed to insert country: %v", err)
	}

	id, err := service.AddTeam(model.CreateTeamDTO{
		Name:           "Test Team",
		CountryID:      country.ID,
		IsNationalTeam: true,
	})
	if err != nil {
		t.Fatalf("failed to add team: %v", err)
	}

	// Now retrieve
	got, err := service.GetTeam(id)
	if err != nil {
		t.Fatalf("failed to get team: %v", err)
	}

	if got.Name != "Test Team" || got.CountryID != country.ID || !got.IsNationalTeam {
		t.Errorf("unexpected team data: %+v", got)
	}
}
