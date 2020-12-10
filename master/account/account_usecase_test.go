package account

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vivaldy22/simpleRestApiLA/tools/queries"

	"github.com/stretchr/testify/assert"
)

func TestNewUseCase(t *testing.T) {
	db, _ := NewMockDB()
	repo := NewRepo(db)
	uc := NewUseCase(repo)
	assert.NotNil(t, uc)
}

func TestAccUseCase_GetByAccNum(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()
	uc := NewUseCase(repo)

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"}).
		AddRow(ma1.AccountNumber, ma1.CustomerNumber, ma1.Balance)

	mock.ExpectQuery(regexp.QuoteMeta(queries.GET_ACCOUNT_BY_ACC_NUM)).WithArgs(ma1.AccountNumber).WillReturnRows(rows)

	data, err := uc.GetByAccNum(ma1.AccountNumber)
	assert.NotNil(t, data)
	assert.NoError(t, err)
}

func TestAccUseCase_GetByAccNumError(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()
	uc := NewUseCase(repo)

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"})

	mock.ExpectQuery(regexp.QuoteMeta(queries.GET_ACCOUNT_BY_ACC_NUM)).WithArgs(ma1.AccountNumber).WillReturnRows(rows)

	data, err := uc.GetByAccNum(ma1.AccountNumber)
	assert.Nil(t, data)
	assert.Error(t, err)
}

/*func TestAccUseCase_TransferBalance(t *testing.T) {
	db, mock := NewMockDB()
	repo := &accRepo{db}
	defer repo.db.Close()
	uc := &accUseCase{repo}

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"}).
		AddRow(ma1.AccountNumber, ma1.CustomerNumber, ma1.Balance).
		AddRow(ma2.AccountNumber, ma2.CustomerNumber, ma2.Balance)

	mock.ExpectQuery(regexp.QuoteMeta(queries.GET_ACCOUNT_BY_ACC_NUM)).WithArgs(ma1.AccountNumber).WillReturnRows(rows)

	msg, err := uc.TransferBalance(ma1.AccountNumber, mt.ToAccountNumber, mt.Amount)
	assert.Equal(t, "Transfer success", msg)
	assert.NoError(t, err)
}*/
