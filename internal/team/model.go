package team

type Country struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"not null"`
	Shortcut   string `gorm:"size:3;not null"`
	Federation string `gorm:"type:text;not null;check:federation IN ('E','AM','AZ','AF')"`

	Teams []Team `gorm:"foreignKey:CountryID"`
}

type Team struct {
	ID             int64  `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"not null"`
	CountryID      int64  `gorm:"not null"`
	IsNationalTeam bool   `gorm:"not null"`

	Country Country `gorm:"foreignKey:CountryID"`
}
