package models

type Transaction struct {
	ID       int64   `gorm:"id" json:"id"`
	Amount   float64 `gorm:"amount" json:"amount"`
	Type     string  `gorm:"amount" json:"type"`
	ParentID int64   `gorm:"parent_id" json:"parent_id"`
}
