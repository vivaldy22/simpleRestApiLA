package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vivaldy22/simpleRestApiLA/tools/viper"
)

func InitDB() *sql.DB {
	dbUser := viper.GetEnv("DB_USER", "root")
	dbPass := viper.GetEnv("DB_PASSWORD", "password")
	dbHost := viper.GetEnv("DB_HOST", "localhost")
	dbPort := viper.GetEnv("DB_PORT", "3306")
	schemaName := viper.GetEnv("DB_SCHEMA", "schema")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, err := sql.Open("mysql", dbPath)
	if err != nil {
		panic(err)
	}
	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}
	return dbConn
}
