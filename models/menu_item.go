package models

type MenuItem struct {
	BaseModel
	MerchantId   uint
	Merchant     MerchantProfile
	Name         string  `json:"name" gorm:"not null"`
	Description  string  `json:"description" gorm:"not null"`
	Rating       float64 `json:"rating" gorm:"default:0;not null"`
	Price        float64 `json:"price" gorm:"default:0;not null"`
	MenuImageURL string  `json:"menu_image_url,omitempty"`
	TotalSold    int     `json:"total_sold" gorm:"default:0;not null"`
	Category     string  `json:"category,omitempty"`
}
