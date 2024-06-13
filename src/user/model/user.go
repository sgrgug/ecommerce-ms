package model

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"size:191;not null;unique"`
}
