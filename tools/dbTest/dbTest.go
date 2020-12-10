package dbTest

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDbTest() *sql.DB {
	database, _ := sql.Open("sqlite3", ":memory:")

	database.Exec("CREATE TABLE IF NOT EXISTS tb_customer (customer_number TEXT PRIMARY KEY, name TEXT)")
	database.Exec("CREATE TABLE IF NOT EXISTS tb_account (account_number TEXT PRIMARY KEY, customer_number TEXT, balance INTEGER)")

	database.Exec("INSERT INTO tb_customer VALUES ('2001', 'Dummy 1')")
	database.Exec("INSERT INTO tb_customer VALUES ('2002', 'Dummy 2'")
	database.Exec("INSERT INTO tb_account VALUES ('555111','2001', 20000)")
	database.Exec("INSERT INTO tb_account VALUES ('555112','2002', 10000)")

	return database
}
