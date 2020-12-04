package queries

const (
	GET_ACCOUNT_BY_ACC_NUM = `SELECT * FROM tb_account WHERE account_number = ?`
	TRANSFER_BALANCE       = `UPDATE tb_account
							  SET balance = CASE account_number
										    WHEN ? THEN balance - ?
										    WHEN ? THEN balance + ?
										    ELSE balance
										    END
							  WHERE account_number IN(?, ?);`
)
