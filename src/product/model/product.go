package model

type Product struct {
	ID          string  `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"size;not null"`
	Price       float64 `gorm:"type:double precision"`
	Quantity    uint32  `gorm:"type:integer"`
	Category    string  `gorm:"size:255;not null"`
}
