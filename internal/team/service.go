package team

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) AddTeam(dto CreateTeamDTO) (Team, error) {
	team := Team{
		Name:           dto.Name,
		CountryID:      dto.CountryID,
		IsNationalTeam: dto.IsNationalTeam,
	}

	if err := s.db.Create(&team).Error; err != nil {
		return Team{}, err
	}
	return team, nil
}

func (s *Service) ListTeams() ([]ShortTeamDTO, error) {
	var teams []Team
	if err := s.db.Find(&teams).Error; err != nil {
		return nil, err
	}

	var result []ShortTeamDTO
	for _, t := range teams {
		result = append(result, toTeamDTO(t))
	}

	return result, nil
}

func toTeamDTO(t Team) ShortTeamDTO {
	return ShortTeamDTO{
		ID:             t.ID,
		Name:           t.Name,
		CountryID:      t.CountryID,
		IsNationalTeam: t.IsNationalTeam,
	}
}
