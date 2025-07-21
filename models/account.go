package models

type AccountType string

const (
	MainBalance AccountType = "main_balance"
	Points      AccountType = "points"
)

type Account struct {
	BaseModel
	Name                 string        `json:"name" gorm:"not null"`
	Balance              float64       `json:"balance" gorm:"default:0"`
	UserId               uint          `json:"user_id" gorm:"not null;index:idx_user_id"`
	AccountType          AccountType   `json:"account_type" gorm:"not null;default:main_balance"`
	User                 User          `json:"-"`
	SentTransactions     []Transaction `json:"sent_transactions,omitempty" gorm:"foreignKey:SenderAccountId"`
	ReceivedTransactions []Transaction `json:"received_transactions,omitempty" gorm:"foreignKey:ReceiverAccountId"`
	QRCodes              []QrCode      `json:"qr_codes,omitempty" gorm:"foreignKey:ReceiverAccountID"`
}
