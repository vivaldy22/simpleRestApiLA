package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vivaldy22/simpleRestApiLA/tools/myViper"
)

func InitDB() (*sql.DB, error) {
	dbUser := myViper.GetEnv("DB_USER", "root")
	dbPass := myViper.GetEnv("DB_PASSWORD", "password")
	dbHost := myViper.GetEnv("DB_HOST", "localhost")
	dbPort := myViper.GetEnv("DB_PORT", "3306")
	schemaName := myViper.GetEnv("DB_SCHEMA", "schema")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, err := sql.Open("mysql", dbPath)
	if err != nil {
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("Database connected, %v\n", schemaName)
	return dbConn, nil
}
