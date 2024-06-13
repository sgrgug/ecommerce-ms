package model

type Order struct {
	ID         string  `gorm:"primaryKey;autoIncrement"`
	User_Id    string  `gorm:"not null"`
	Product_Id string  `gorm:"not null"`
	Quantity   uint32  `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	Status     string  `gorm:"not null;default:'pending'"`
}
