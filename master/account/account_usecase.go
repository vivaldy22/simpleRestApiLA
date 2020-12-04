package account

import (
	"github.com/vivaldy22/simpleRestApiLA/models"
)

type accUseCase struct {
	accRepo models.AccountRepo
}

func (a *accUseCase) GetByAccNum(accNum string) (*models.Account, error) {
	res, err := a.accRepo.GetByAccNum(accNum)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *accUseCase) Transfer(from, to string, amount int) error {
	err := a.accRepo.Transfer(from, to, amount)
	if err != nil {
		return err
	}
	return nil
}

func NewUseCase(accRepo models.AccountRepo) models.AccountUseCase {
	return &accUseCase{accRepo}
}
