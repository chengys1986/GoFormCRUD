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
	// Global
	err  error
	conn *sqlx.DB
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
	conn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return conn, err
}

// Close ...
func CloseDB(c io.Closer) {
	fmt.Println("DB Closing")

	if err := c.Close(); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}

func GetStudents() []model.Student {
	fmt.Println("DBService: Get Students")

	students := []model.Student{}
	query := "SELECT * FROM Student;"
	err = conn.Select(&students, query)
	if err != nil {
		fmt.Println("Get Student error:", err)
		return nil
	}
	fmt.Printf("%#v\n", students)
	fmt.Println("load students successfully.")
	return students
}

func NewStudent(student model.Student) bool {
	fmt.Println("DBService: New a student")
	var res bool = true

	query := "INSERT INTO Student"
	query += "(Name, Description)"
	query += "VALUES (:Name, :Description)"
	_, err = conn.NamedExec(query, student)
	if err != nil {
		res = false
	}
	return res
}

func GetStudent(id string) (error, model.Student) {
	fmt.Println("DBService: Get Student")
	student := model.Student{}

	query := "SELECT * FROM Student "
	query += "WHERE studentid = "
	query += id
	query += ";"
	err = conn.Get(&student, query)

	return err, student
}

func UpdateStudent(student model.Student) bool {
	fmt.Println("DBService: Update the student")
	var res bool = true

	query := "UPDATE Student "
	query += "SET Name = :Name, Description = :Description "
	query += "WHERE StudentID = :StudentID ;"

	_, err = conn.NamedExec(query, student)
	if err != nil {
		res = false
	}
	return res
}

func DeleteStudent(id string) bool {
	fmt.Println("DBService: Delete Student")
	var res bool = true

	intId, err := strconv.Atoi(id)
	query := "DELETE FROM Student "
	query += "WHERE StudentID = :StudentID ;"
	student := model.Student{}
	student.StudentID = intId

	_, err = conn.NamedExec(query, student)
	if err != nil {
		res = false
	}

	return res
}
