package account

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/vivaldy22/simpleRestApiLA/tools/queries"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vivaldy22/simpleRestApiLA/models"
)

var ma1 = &models.Account{
	AccountNumber:  "123456",
	CustomerNumber: "2001",
	Balance:        "20000",
}

var ma2 = &models.Account{
	AccountNumber:  "123457",
	CustomerNumber: "2002",
	Balance:        "10000",
}

var mt = &models.Transfer{
	ToAccountNumber: "123457",
	Amount:          "100",
}

func NewMockDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occurred when opening a mock database connection")
	}
	return db, mock
}

func TestNewRepo(t *testing.T) {
	db, _ := NewMockDB()
	repo := NewRepo(db)
	assert.NotNil(t, repo)
}

func TestAccRepo_GetByAccNum(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"}).
		AddRow(ma1.AccountNumber, ma1.CustomerNumber, ma1.Balance)

	mock.ExpectQuery(regexp.QuoteMeta(queries.GET_ACCOUNT_BY_ACC_NUM)).WithArgs(ma1.AccountNumber).WillReturnRows(rows)

	data, err := repo.GetByAccNum(ma1.AccountNumber)
	assert.NotNil(t, data)
	assert.NoError(t, err)
}

func TestAccRepo_GetByAccNumError(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"})

	mock.ExpectQuery(regexp.QuoteMeta(queries.GET_ACCOUNT_BY_ACC_NUM)).WithArgs(ma1.AccountNumber).WillReturnRows(rows)

	data, err := repo.GetByAccNum(ma1.AccountNumber)
	assert.Nil(t, data)
	assert.Error(t, err)
}

func TestAccRepo_Transfer(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(regexp.QuoteMeta(queries.TRANSFER_BALANCE))
	prep.ExpectExec().WithArgs(ma1.AccountNumber, mt.Amount, mt.ToAccountNumber, mt.Amount, ma1.AccountNumber, mt.ToAccountNumber).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.Transfer(ma1.AccountNumber, mt.ToAccountNumber, mt.Amount)
	assert.NoError(t, err)
}

func TestAccRepo_TransferError(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(regexp.QuoteMeta(queries.TRANSFER_BALANCE))
	prep.ExpectExec().WithArgs(ma1.AccountNumber, mt.Amount, mt.ToAccountNumber, mt.Amount, ma1.AccountNumber, mt.ToAccountNumber).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.Transfer(ma1.AccountNumber, mt.ToAccountNumber, mt.Amount)
	assert.Error(t, err)
}
