package models

type User struct {
	BaseModel
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"-"`
	Accounts []Account `json:"accounts,omitempty"` // we dont have to put gorm fk here because we haev UserId at account, so gorm will assume it is the fk

	// contacts i created
	Contacts []Contact `json:"contacts,omitempty" gorm:"foreignKey:OwnerID"` // the reason we put it here foreignkey ownerId is because this is an user struct and at contact, we have it as owner, not UserId, so gorm needs precise fk explicitly.
}
