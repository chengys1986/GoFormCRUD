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
		"exp",       // Username
		"exp",       // Password
		"localhost", // Host
		"3306",      // Port
		"exp",       // Name
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
