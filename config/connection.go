package config

import (
	"database/sql"
	"fmt"
	"github.com/vivaldy22/simpleRestApiLA/tools/viper"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dbUser := viper.ViperGetEnv("DB_USER", "root")
	dbPass := viper.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := viper.ViperGetEnv("DB_HOST", "localhost")
	dbPort := viper.ViperGetEnv("DB_PORT", "3306")
	schemaName := viper.ViperGetEnv("DB_SCHEMA", "schema")

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
