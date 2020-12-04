package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vivaldy22/simpleRestApiLA/tools/viper"
)

func InitDB() (*sql.DB, error) {
	dbUser := viper.GetEnv("DB_USER", "root")
	dbPass := viper.GetEnv("DB_PASSWORD", "password")
	dbHost := viper.GetEnv("DB_HOST", "localhost")
	dbPort := viper.GetEnv("DB_PORT", "3306")
	schemaName := viper.GetEnv("DB_SCHEMA", "schema")
	log.Println(dbUser, dbPass, dbHost, dbPort, schemaName)

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
