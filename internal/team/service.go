package team

type Service struct {
	teams []Team
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddTeam(name string) Team {
	team := Team{
		ID:   len(s.teams) + 1,
		Name: name,
	}
	s.teams = append(s.teams, team)
	return team
}

func (s *Service) ListTeams() []Team {
	return s.teams
}
