package models

type Account struct {
	AccountNumber string `json:"account_number"`
	CustomerNumber string `json:"customer_number"`
	Balance int `json:"balance"`
}

type AccountRepo interface {
	GetByAccNum(accNum string) (*Account, error)
	Transfer(from, to string, amount int) error
}

type AccountUseCase interface {
	GetByAccNum(accNum string) (*Account, error)
	Transfer(from, to string, amount int) error
}