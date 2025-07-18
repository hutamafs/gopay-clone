package models

type Contact struct {
	BaseModel
	OwnerID  uint `json:"owner_id" gorm:"not null"`
	Owner    User `gorm:"foreignKey:OwnerID"`
	TargetID uint `json:"target_id" gorm:"not null"`
	Target   User `gorm:"foreignKey:TargetID"`
}
