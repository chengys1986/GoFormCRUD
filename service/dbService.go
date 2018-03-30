package service

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
	"../model"
)

var (
	err error

	// DB ...
	DB *sqlx.DB

	// Data Source Name
	dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tx_isolation='READ-COMMITTED'",
		"exp",       // Username
		"exp",       // Password
		"localhost", // Host
		"3306",      // Port
		"exp",       // Name
	)
)

// Load ...
func loadDB() {
	fmt.Println("DB Loading")

	// Start database connection
	DB, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		closeDB(DB)
		return
	}
	//	defer closeDB(DB)
}

// Close ...
func closeDB(c io.Closer) {
	fmt.Println("DB Closing")

	if err := c.Close(); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}

func GetStudents() []model.Student {
	fmt.Println("DBService: Get Students")

	loadDB()

	students := []model.Student{}
	query := "SELECT * FROM Student;"
	err = DB.Select(&students, query)
	if err != nil {
		fmt.Println("Get Student error:", err)
		closeDB(DB)
		return nil
	}
	fmt.Printf("%#v\n", students)
	fmt.Println("load students successfully.")
	closeDB(DB)
	return students
}
