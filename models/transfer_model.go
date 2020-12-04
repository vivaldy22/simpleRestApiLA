package models

type Transfer struct {
	ToAccountNumber string `json:"to_account_number"`
	Amount int `json:"amount"`
}
