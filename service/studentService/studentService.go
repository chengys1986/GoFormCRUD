package studentService

import (
	"fmt"
	"strconv"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

import (
	"../../model"
	"../db"
)

func GetStudents() []model.Student {
	fmt.Println("DBService: Get Students")

	students := []model.Student{}
	query := "SELECT * FROM Student;"
	err := db.Conn.Select(&students, query)
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
	_, err := db.Conn.NamedExec(query, student)
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
	err := db.Conn.Get(&student, query)

	return err, student
}

func UpdateStudent(student model.Student) bool {
	fmt.Println("DBService: Update the student")
	var res bool = true

	query := "UPDATE Student "
	query += "SET Name = :Name, Description = :Description "
	query += "WHERE StudentID = :StudentID ;"

	_, err := db.Conn.NamedExec(query, student)
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

	_, err = db.Conn.NamedExec(query, student)
	if err != nil {
		res = false
	}

	return res
}
