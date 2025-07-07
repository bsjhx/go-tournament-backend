package model

type CreateTeamDTO struct {
	Name           string `json:"name" binding:"required"`
	CountryID      int64  `json:"country_id" binding:"required"`
	IsNationalTeam bool   `json:"is_national_team"`
}

type ShortTeamDTO struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	CountryID      int64  `json:"country_id"`
	IsNationalTeam bool   `json:"is_national_team"`
}

type ShortCountryDTO struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Shortcut   string `json:"shortcut"`
	Federation string `json:"federation"`
}
