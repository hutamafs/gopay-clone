package models

type MenuCategory string

const (
	MainCourse MenuCategory = "main_course"
	Appetizer  MenuCategory = "appetizer"
	Dessert    MenuCategory = "dessert"
	Beverage   MenuCategory = "beverage"
	Snack      MenuCategory = "snack"
	Drink      MenuCategory = "drink"
)

type MenuItem struct {
	BaseModel
	MerchantId   uint            `json:"merchant_id" gorm:"not null"`
	Merchant     MerchantProfile `json:"merchant" gorm:"foreignKey:MerchantId"`
	Name         string          `json:"name" gorm:"not null"`
	Description  string          `json:"description" gorm:"not null"`
	Rating       float64         `json:"rating" gorm:"default:0;not null"`
	Price        float64         `json:"price" gorm:"default:0;not null"`
	MenuImageURL string          `json:"menu_image_url,omitempty"`
	TotalSold    int             `json:"total_sold" gorm:"default:0;not null"`
	Category     MenuCategory    `json:"category" gorm:"default:main_course"`
	IsAvailable  bool            `json:"is_available" gorm:"default:true"`
}
