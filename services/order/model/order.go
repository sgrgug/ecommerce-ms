package model

type Order struct {
	ID         uint    `gorm:"primaryKey"`
	User_Id    uint    `gorm:"not null"`
	Product_Id uint    `gorm:"not null"`
	Quantity   int64   `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	Status     string  `gorm:"not null;default:'pending'"`
}
