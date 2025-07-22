package models

type OrderItem struct {
	BaseModel
	OrderID    uint     `json:"order_id" gorm:"not null"`
	MenuItemID uint     `json:"menu_item_id" gorm:"not null"`
	Order      Order    `json:"-" gorm:"foreignKey:OrderID"`
	MenuItem   MenuItem `json:"-" gorm:"foreignKey:MenuItemID"`
	Quantity   int      `json:"quantity" gorm:"not null"`
	Price      float64  `json:"price" gorm:"not null"` // price at time of order
	Notes      string   `json:"notes"`
}
