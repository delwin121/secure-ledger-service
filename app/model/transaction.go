package model

type Transaction struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}
