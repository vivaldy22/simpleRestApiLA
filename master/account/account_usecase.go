package account

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/vivaldy22/simpleRestApiLA/models"
)

type accUseCase struct {
	accRepo models.AccountRepo
}

func (a *accUseCase) GetByAccNum(accNum string) (*models.Account, error) {
	return a.accRepo.GetByAccNum(accNum)
}

func (a *accUseCase) Transfer(from, to, amount string) (string, error) {
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		return "Amount is not a valid number", err
	}

	accFrom, err := a.accRepo.GetByAccNum(from)
	if err != nil {
		return fmt.Sprintf("Account %v doesn't exist", from), err
	}

	intBalance, err := strconv.Atoi(accFrom.Balance)
	if err != nil {
		return "Balance is not a valid number", err
	}

	if intAmount == 0 {
		return "You need to fill the amount, minimum 1", errors.New("amount can not be 0")
	}

	_, err = a.accRepo.GetByAccNum(to)
	if err != nil {
		return fmt.Sprintf("Account %v doesn't exist", to), err
	}

	if intBalance-intAmount < 0 {
		return "Your Balance is not enough", errors.New("minus balance")
	}

	if from == to {
		return "You can not transfer to your own account", errors.New("can not transfer same account")
	}

	if err = a.accRepo.Transfer(from, to, amount); err != nil {
		return "Transfer failed, see logs", err
	}

	return "Transfer success", nil

}

func NewUseCase(accRepo models.AccountRepo) models.AccountUseCase {
	return &accUseCase{accRepo}
}
