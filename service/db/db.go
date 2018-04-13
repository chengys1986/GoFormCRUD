package db

import (
	"fmt"
	"io"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

import (
	"../config"
)

var (
	// Global
	Conn *sqlx.DB

	err error
)

// Load ...
func InitDB() (*sqlx.DB, error) {
	fmt.Println("DB init")

	// Data Source Name
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?tx_isolation='READ-COMMITTED'",
		config.ReadConfig("Database.Username"), // Username
		config.ReadConfig("Database.Password"), // Password
		config.ReadConfig("Database.Host"),     // Host
		config.ReadConfig("Database.Port"),     // Port
		config.ReadConfig("Database.Name"),     // Name
	)
	Conn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return Conn, err
}

// Close ...
func CloseDB(c io.Closer) {
	fmt.Println("DB Closing")

	if err = c.Close(); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
