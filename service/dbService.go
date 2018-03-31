package service

import (
	"fmt"
	"io"
	"log"
	"strconv"
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

func NewStudent(student model.Student) bool {
	fmt.Println("DBService: New a student")
	var res bool = true
	loadDB()

	query := "INSERT INTO Student"
	query += "(Name, Description)"
	query += "VALUES (:Name, :Description)"
	_, err = DB.NamedExec(query, student)
	if err != nil {
		res = false
	}
	closeDB(DB)
	return res
}

func GetStudent(id string) (error, model.Student) {
	fmt.Println("DBService: Get Student")
	student := model.Student{}
	loadDB()

	query := "SELECT * FROM Student "
	query += "WHERE studentid = "
	query += id
	query += ";"
	err = DB.Get(&student, query)

	closeDB(DB)
	return err, student
}

func UpdateStudent(student model.Student) bool {
	fmt.Println("DBService: Update the student")
	var res bool = true
	loadDB()

	query := "UPDATE Student "
	query += "SET Name = :Name, Description = :Description "
	query += "WHERE StudentID = :StudentID ;"

	_, err = DB.NamedExec(query, student)
	if err != nil {
		res = false
	}
	closeDB(DB)
	return res
}

func DeleteStudent(id string) bool {
	fmt.Println("DBService: Delete Student")
	var res bool = true
	loadDB()

	query := "DELETE FROM Student "
	query += "WHERE StudentID = :StudentID ;"
	intId, err := strconv.Atoi(id)
	student := model.Student{}
	student.StudentID = intId
	_, err = DB.NamedExec(query, student)
	if err != nil {
		res = false
	}

	closeDB(DB)
	return res
}
