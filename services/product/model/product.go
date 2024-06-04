package model

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"size;not null"`
	Price       float64 `gorm:"type:double precision"`
	Quantity    int64   `gorm:"type:integer"`
	Category    string  `gorm:"size:255;not null"`
}
