package models

type Orders struct {
	ID      uint `gorm:"primaryKey" json:"orders_id"`
	ItemsID uint `json:"items_id"`
	// Item     *Item
	Quantity uint `json:"quantity"`
}
