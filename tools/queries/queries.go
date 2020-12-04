package queries

const (
	GET_ACCOUNT_BY_ACC_NUM = `SELECT * FROM tb_account WHERE account_number = ?`
	TRANSFER_BALANCE = `UPDATE tb_account
						SET balance = balance - ?
						WHERE account_number = ?;
						UPDATE tb_account
						SET balance = balance + ?
						WHERE account_number = ?;
`
)
