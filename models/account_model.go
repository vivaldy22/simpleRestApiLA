package models

type Account struct {
	AccountNumber  string `json:"account_number"`
	CustomerNumber string `json:"customer_number"`
	Balance        string `json:"balance"`
}

type AccountRepo interface {
	GetByAccNum(accNum string) (*Account, error)
	Transfer(from, to, amount string) error
}

type AccountUseCase interface {
	GetByAccNum(accNum string) (*Account, error)
	TransferBalance(from, to, amount string) (string, error)
}
