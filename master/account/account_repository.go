package account

import (
	"database/sql"

	"github.com/vivaldy22/simpleRestApiLA/models"
	"github.com/vivaldy22/simpleRestApiLA/tools/queries"
)

type accRepo struct {
	*sql.DB
}

func (a *accRepo) GetByAccNum(accNum string) (*models.Account, error) {
	var acc = new(models.Account)
	row := a.QueryRow(queries.GET_ACCOUNT_BY_ACC_NUM, accNum)
	err := row.Scan(&acc.AccountNumber, &acc.CustomerNumber, &acc.Balance)

	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (a *accRepo) Transfer(from, to, amount string) error {
	tx, err := a.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.TRANSFER_BALANCE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(from, amount, to, amount, from, to)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func NewRepo(db *sql.DB) models.AccountRepo {
	return &accRepo{db}
}
