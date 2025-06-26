package team

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	db.AutoMigrate(&Team{})
	return &Service{db: db}
}

func (s *Service) AddTeam(name string) Team {
	team := Team{Name: name}
	s.db.Create(&team)
	return team
}

func (s *Service) ListTeams() []Team {
	var teams []Team
	s.db.Find(&teams)
	return teams
}
