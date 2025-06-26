package team

type Team struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
