package models

type Account struct {
	AccountNumber string `json:"account_number"`
	CustomerNumber string `json:"customer_number"`
	Balance int `json:"balance"`
}

type AccountRepo interface {
	GetByAccNum(accNum string) (Account, error)
	Transfer(toAccNum string, amount int) error
}
