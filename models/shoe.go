package models

type Shoe struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Size        int     `json:"size"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
