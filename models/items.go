package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"items_id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}
