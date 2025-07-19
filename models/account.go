package models

type Account struct {
	BaseModel
	Name                 string        `json:"name" gorm:"not null"`
	Balance              float64       `json:"balance" gorm:"default:0"`
	UserId               uint          `json:"user_id" gorm:"not null"`
	User                 User          `json:"-"`
	SentTransactions     []Transaction `json:"sent_transactions,omitempty" gorm:"foreignKey:SenderAccountId"`
	ReceivedTransactions []Transaction `json:"received_transactions,omitempty" gorm:"foreignKey:ReceiverAccountId"`
	QRCodes              []QrCode      `json:"qr_codes,omitempty" gorm:"foreignKey:ReceiverAccountID"`
}
