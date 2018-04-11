package main

import (
	"fmt"
	"log"
	"net/http"
)

import (
	"./controller/home"
	"./controller/student"
	"./service/db"
)

func main() {
	fmt.Println("Starting up")

	// init db
	conn, err := db.InitDB()
	if err != nil {
		fmt.Printf("DB Error: $s\n", err.Error())
		return
	}
	defer db.CloseDB(conn)

	// Routing
	http.HandleFunc("/", homeController.HomeController)
	http.HandleFunc("/student/new", studentController.StudentNewController)
	http.HandleFunc("/student/list", studentController.StudentListController)
	http.HandleFunc("/student/form", studentController.StudentFormController)
	http.HandleFunc("/student/edit", studentController.StudentEditController)
	http.HandleFunc("/student/delete", studentController.StudentDeleteController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
