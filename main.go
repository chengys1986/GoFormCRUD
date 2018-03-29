package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
)

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// DB ...
	DB *sqlx.DB
)

func main() {
	var err error

	// Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tx_isolation='READ-COMMITTED'",
		"exp",       // Username
		"exp",       // Password
		"localhost", // Host
		"3306",      // Port
		"exp",       //Name
	)

	// Start database connection
	DB, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	defer Close(DB)

	// Routing
	http.HandleFunc("/student/new", studentNewController)
	http.HandleFunc("/student/form", studentFormController)
	http.HandleFunc("/student/list", studentListController)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Close ...
func Close(c io.Closer) {
	// Note: do we need to add recover() here?
	if err := c.Close(); err != nil {
		log.Printf(err.Error())
	}
}
func studentListController(w http.ResponseWriter, r *http.Request) {
	var err error
	students := []Student{}

	query := "SELECT * FROM Student;"

	err = DB.Select(&students, query)

	if err != nil {
		fmt.Println("error!", err)
	} else {

		fmt.Printf("%#v", students)

		w.Header().Set("Content-Type", "text/html; charset-utf-8")

		var tplParam = struct {
			Type        string
			StudentList []Student
		}{
			"LIST",
			students,
		}
		tplPath := "template/student.list.tpl.html"
		tpl := template.New(path.Base(tplPath))

		var outStr string

		if tpl, err := tpl.ParseFiles(tplPath); err != nil {
			log.Println(err)
		} else {
			var doc bytes.Buffer
			tpl.Execute(&doc, tplParam)
			outStr = doc.String()
		}
		w.Write([]byte(outStr))
	}
}

type Student struct {
	StudentID   int    `json:"studentid" db:"StudentID"`
	Name        string `json:"name" db:"Name"`
	Description string `json:"description" db:"Description"`
}

func studentFormController(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "form %s", r.URL.Path[1:])
	var err error

	w.Header().Set("Content-Type", "text/html; charset-utf-8")

	student := Student{}

	query := "SELECT * FROM Student "
	query += "WHERE studentid = 1;"

	err = DB.Get(&student, query)

	if err != nil {
		fmt.Println("error!", err)
	} else {

		fmt.Printf("%#v", student)

		var tplParam = struct {
			Type        string
			Name        string
			Description string
		}{
			"FORM",
			student.Name,
			student.Description,
		}
		tplPath := "template/student.form.tpl.html"
		tpl := template.New(path.Base(tplPath))

		var outStr string

		if tpl, err := tpl.ParseFiles(tplPath); err != nil {
			log.Println(err)
		} else {
			var doc bytes.Buffer
			tpl.Execute(&doc, tplParam)
			outStr = doc.String()
		}
		w.Write([]byte(outStr))
	}
}
func studentNewController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset-utf-9")
		var tplParam = struct {
			Type string
		}{
			"NEW",
		}
		tplPath := "template/student.new.tpl.html"
		tpl := template.New(path.Base(tplPath))

		var outStr string

		if tpl, err := tpl.ParseFiles(tplPath); err != nil {
			log.Println(err)
		} else {
			var doc bytes.Buffer
			tpl.Execute(&doc, tplParam)
			outStr = doc.String()
		}
		w.Write([]byte(outStr))
	} else {
		var err error

		query := "INSERT INTO Student"
		query += "(Name, Description)"
		query += "VALUES (:Name, :Description)"

		_, err = DB.NamedExec(query, map[string]interface{}{
			"Name":        r.FormValue("name"),
			"Description": r.FormValue("description"),
		})

		if err != nil {
			fmt.Println("error!", err)
		}
	}
}
