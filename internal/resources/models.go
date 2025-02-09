package resources

type Transaction struct {
	ID             string  `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Description    string  `json:"description"`
	Date           string  `json:"transaction_date"`
	PurchaseAmount float64 `json:"purchase_amount"`
}

type ConvertedTransaction struct {
	Transaction
	ExchangeRate    float64 `json:"exchange_rate"`
	ConvertedAmount float64 `json:"converted_amount"`
	Currency        string  `json:"currency"`
}

type Error struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
}
