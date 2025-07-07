package services

import (
	"errors"
	"github.com/bsjhx/tournament-backend/internal/team/model"
	"gorm.io/gorm"
	"log"
)

type TeamService struct {
	db *gorm.DB
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{db: db}
}

func (s *TeamService) AddTeam(dto model.CreateTeamDTO) (int64, error) {
	newTeam := model.Team{
		Name:           dto.Name,
		CountryID:      dto.CountryID,
		IsNationalTeam: dto.IsNationalTeam,
	}

	if err := s.db.Create(&newTeam).Error; err != nil {
		log.Fatalln(errors.New("failed to create team"))
		return -1, err
	}
	return newTeam.ID, nil
}

func (s *TeamService) GetTeam(id int64) (model.ShortTeamDTO, error) {
	var foundTeam model.Team
	if err := s.db.First(&foundTeam, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ShortTeamDTO{}, nil
		}
		return model.ShortTeamDTO{}, err
	}

	return toTeamDTO(foundTeam), nil
}

func (s *TeamService) ListTeams() ([]model.ShortTeamDTO, error) {
	var teams []model.Team
	if err := s.db.Find(&teams).Error; err != nil {
		return nil, err
	}

	var result []model.ShortTeamDTO
	for _, t := range teams {
		result = append(result, toTeamDTO(t))
	}

	return result, nil
}

func toTeamDTO(t model.Team) model.ShortTeamDTO {
	return model.ShortTeamDTO{
		ID:             t.ID,
		Name:           t.Name,
		CountryID:      t.CountryID,
		IsNationalTeam: t.IsNationalTeam,
	}
}
