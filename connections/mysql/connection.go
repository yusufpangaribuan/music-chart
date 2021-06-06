package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var mysqlClient *sql.DB

func Initialize() {
	connectDB()
}

func connectDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlClient, err = sql.Open(os.Getenv("SQL_DRIVER"), os.Getenv("SQL_DOMAIN"))
	if err != nil {
		log.Panicln("Fail to connect DB", err)
	}
	err = mysqlClient.Ping()
	if err != nil {
		log.Panicln("Fail to ping", err)
	}
	mysqlClient.SetMaxIdleConns(20)
	mysqlClient.SetMaxOpenConns(100)
}

// GetSQLClient to get SQL Client
func GetSQLClient() *sql.DB {
	return mysqlClient
}
