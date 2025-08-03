package models

// Transaction represents a financial transaction
type Transaction struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
}

// TableName specifies the table name for the Transaction model
func (Transaction) TableName() string {
	return "transactions"
} 